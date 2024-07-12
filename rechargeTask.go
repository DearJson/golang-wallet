package main

import (
	"context"
	"fmt"
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
)

// 归集任务
func main() {
	if g.Cfg().GetBool("bsc.address_recharge") {
		fmt.Println("开始了")
		bscRecharge()
	}
}

// bscRecharge 币安链任务归集
func bscRecharge() {
	//判断一下是否有配置手续费地址和归集地址
	bnbMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.bnbMergeAddress")
	bnbFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.bnbFeeAddress")
	bnbFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.bnbFeeAddressPrivateKey")
	//如果未配置出金地址，退出
	if bnbMergeAddress.ConfigValue == "" || bnbFeeAddress.ConfigValue == "" || bnbFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("归集地址或手续费私钥地址未配置，退出归集")
		return
	}
	bnbGasLimit, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	bnbGasPrice, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")

	bnbFeePrivateKey, _ := library.DecryptByAes(bnbFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "bsc").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	minFee := gconv.Int64(bnbGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(bnbGasLimit.ConfigValue)

	for _, value := range list {

		//查询当前地址余额
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)

		var last int64
		if value.ContractAddress == "0x1000000000000000000000000000000000000000" {
			if balance.Int64()-minFee-gconv.Int64(gconv.Float64(value.Amount)*math.Pow10(18)) < 0 {
				last = balance.Int64() + minFee - gconv.Int64(gconv.Float64(value.Amount)*math.Pow10(18))
			} else {
				last = 0
			}
		} else {
			if balance.Int64()-minFee < 0 {
				last = minFee - balance.Int64()
			} else {
				last = 0
			}
		}
		//需要先转手续费
		if last > 0 {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "bsc").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "bsc").OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.TransferBnb(string(bnbFeePrivateKey), big.NewInt(last), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount := gconv.Float64(last) / math.Pow10(18)
					g.Model("fee_list").Data(g.Map{"main_chain": "bsc", "coin_name": "bnb", "address": value.ToAddress,
						"amount": amount, "hash": hashResult, "nonce": nonce, "recharge_id": value.Id}).Insert()
					g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 5}).Update()
				}
			}
		} else {
			var (
				hashResult interface{}
				nonce      uint64
				address    *model.Address
				currency   *model.Currency
			)
			//查询币种
			g.Model("currency").Where("main_chain", "bsc").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			//查询当前地址
			g.Model("address").Where("main_chain", "bsc").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)
			//处理金额
			amount := big.NewFloat(gconv.Float64(value.Amount))
			tenDecimal := big.NewFloat(math.Pow(10, float64(currency.Decimals)))
			td := big.NewInt(1000)
			convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
			convertAmount = new(big.Int).Sub(convertAmount, td)

			g.Log().Println("开始操作3\v", convertAmount)

			//查询当前归集地址的最大Nonce
			MaxNonce, _ := g.Model("recharge").Where("main_chain", "bsc").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")
			fmt.Printf("最大的nonce, %v", MaxNonce)
			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
				hashResult, nonce, _ = rpc.TransferBnb(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
			} else {
				hashResult, nonce, _ = rpc.TransferToken(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
			}

			if hashResult != nil {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
			}
		}
	}
}
