package main

import (
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
)

func main() {
	tronMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.tronMergeAddress")
	tronFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.tronFeeAddress")
	tronFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.tronFeeAddressPrivateKey")
	//如果未配置出金地址，退出
	if tronMergeAddress.ConfigValue == "" || tronFeeAddress.ConfigValue == "" || tronFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("波场归集地址或手续费私钥地址未配置，退出归集")
		return
	}

	tronGasLimit, _ := service.SysConfig.GetConfigByKey("sys.tronFee")

	tronFeePrivateKey, _ := library.DecryptByAes(tronFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "tron").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	var cache = service.Cache.New()
	var account *core.Account
	tronGasFee := gconv.Int64(tronGasLimit.ConfigValue) * 1000000
	tronClient, err := rpc.NewClient(gconv.String(cache.Get("tron_rpc_url")))
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("波场节点链接失败，%v", err)
		return
	}

	for _, value := range list {

		//查询当前地址余额
		account, err = tronClient.GetTrxBalance(value.ToAddress)
		if err != nil {
			g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询账号信息 %v 失败 %v", value.ToAddress, err.Error() == "account not found")
			//判断一下，如果是因为账户未激活，使用提币钱包转0.01个trx过去
			if err.Error() == "account not found" {
				_, _ = tronClient.TransferTrx(string(tronFeePrivateKey), tronFeeAddress.ConfigValue, value.ToAddress, decimal.NewFromFloat(0.000001), "")
			}
			continue
		}
		balance := gconv.Int64(account.Balance)

		var last int64
		if value.ContractAddress == "TBRop8PopYu8atWWez3g3ueVtSpseW78b6" {
			if balance-tronGasFee-gconv.Int64(gconv.Float64(value.Amount)*math.Pow10(6)) < 0 {
				last = tronGasFee + gconv.Int64(gconv.Float64(value.Amount)*math.Pow10(6)) - balance
			} else {
				last = 0
			}
		} else {
			if balance-tronGasFee < 0 {
				last = tronGasFee - balance
			} else {
				last = 0
			}
		}
		var txId string

		//需要先转手续费
		if last > 0 {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "tron").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				decimals := decimal.NewFromFloat(math.Pow10(6))
				feeAmount := decimal.NewFromInt(last).Div(decimals)
				feeAmounts, _ := feeAmount.Float64()
				txId, err = tronClient.TransferTrx(string(tronFeePrivateKey), tronFeeAddress.ConfigValue, value.ToAddress, feeAmount, "")

				if err == nil {
					g.Model("fee_list").Data(g.Map{"main_chain": "tron", "coin_name": "trx", "withdraw_address": tronFeeAddress.ConfigValue, "address": value.ToAddress,
						"amount": feeAmounts, "hash": txId[2:], "recharge_id": value.Id}).Insert()
					g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 5}).Update()
				}
			}
		} else {
			var (
				address  *model.Address
				currency *model.Currency
			)
			//查询币种
			g.Model("currency").Where("main_chain", "tron").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			//查询当前地址
			g.Model("address").Where("main_chain", "tron").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)

			if currency.ContractAddress == "TBRop8PopYu8atWWez3g3ueVtSpseW78b6" {
				amd, _ := decimal.NewFromString(value.Amount)
				txId, err = tronClient.TransferTrx(string(privateKey), value.ToAddress, tronMergeAddress.ConfigValue, amd, "")
			} else {
				//处理金额
				amount, _ := decimal.NewFromString(value.Amount)
				tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
				convertAmount := amount.Mul(tenDecimal).BigInt()
				txId, err = tronClient.TransferContract(string(privateKey), value.ToAddress, tronMergeAddress.ConfigValue, value.ContractAddress, convertAmount, tronGasFee)
			}

			if txId != "" {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "imputation_hash": txId[2:]}).Update()
			}
		}
	}

}
