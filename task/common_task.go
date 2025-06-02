package task

import (
	"context"
	"fmt"
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// 归集任务
func rechargeTask() {
	if g.Cfg().GetBool("bsc.address_recharge") || g.Cfg().GetBool("bsc.contract_recharge") {
		bscRecharge()
	}
	if g.Cfg().GetBool("eth.address_recharge") || g.Cfg().GetBool("eth.contract_recharge") {
		ethRecharge()
	}
	if g.Cfg().GetBool("tron.address_recharge") || g.Cfg().GetBool("tron.contract_recharge") {
		tronRecharge()
	}
	if g.Cfg().GetBool("heco.address_recharge") || g.Cfg().GetBool("heco.contract_recharge") {
		hecoRecharge()
	}
	if g.Cfg().GetBool("wemix.address_recharge") {
		wemixRecharge()
	}
	if g.Cfg().GetBool("nac.address_recharge") {
		nacRecharge()
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
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "bsc").Where("withdraw_address", bnbFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.TransferBnb(string(bnbFeePrivateKey), big.NewInt(last), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount := gconv.Float64(last) / math.Pow10(18)
					g.Model("fee_list").Data(g.Map{"main_chain": "bsc", "coin_name": "bnb", "withdraw_address": bnbFeeAddress.ConfigValue, "address": value.ToAddress,
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
			amount, _ := decimal.NewFromString(value.Amount)
			tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
			convertAmount := amount.Mul(tenDecimal).BigInt()

			//查询当前归集地址的最大Nonce
			MaxNonce, _ := g.Model("recharge").Where("main_chain", "bsc").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")
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

// ethRecharge ETH连任务归集
func ethRecharge() {
	g.Log().File("merge_recharge.{Y-m-d}.log").Println("开始ETH归集")
	//判断一下是否有配置手续费地址和归集地址
	ethMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.ethMergeAddress")
	ethFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.ethFeeAddress")
	ethFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.ethFeeAddressPrivateKey")
	//如果未配置出金地址，退出
	if ethMergeAddress.ConfigValue == "" || ethFeeAddress.ConfigValue == "" || ethFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("归集地址或手续费私钥地址未配置，退出归集")
		return
	}
	ethGasLimit, _ := service.SysConfig.GetConfigByKey("sys.ethGasLimit")
	ethGasPrice, _ := service.SysConfig.GetConfigByKey("sys.ethGasPrice")

	ethFeePrivateKey, _ := library.DecryptByAes(ethFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "eth").Where("status", 1).Scan(&list)

	g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询数据条数为 %v", len(list))
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("eth_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	minFee := gconv.Int64(ethGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(ethGasLimit.ConfigValue)

	for _, value := range list {

		//查询当前地址余额
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)

		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询0 %v 余额为 %v", value.ToAddress, balance)

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

		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询1 %v 余额为 %v", value.ToAddress, balance)

		//需要先转手续费
		if last > 0 {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "eth").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "eth").Where("withdraw_address", ethFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询4 %v 余额为 %v", value.ToAddress, balance)
				hashResult, nonce, _ := rpc.TransferEth(string(ethFeePrivateKey), big.NewInt(last), value.ToAddress, gconv.Uint64(MaxNonce))
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询65 %v 余额为 %v", value.ToAddress, balance)
				if hashResult != nil {
					amount := gconv.Float64(last) / math.Pow10(18)
					g.Model("fee_list").Data(g.Map{"main_chain": "eth", "coin_name": "eth", "withdraw_address": ethFeeAddress.ConfigValue, "address": value.ToAddress,
						"amount": amount, "hash": hashResult, "nonce": nonce, "recharge_id": value.Id}).Insert()
					g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 5}).Update()
				} else {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询66 %v 余额为 %v", value.ToAddress, balance)
				}
			} else {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询3 %v 余额为 %v", value.ToAddress, balance)
			}
		} else {
			var (
				hashResult interface{}
				nonce      uint64
				address    *model.Address
				currency   *model.Currency
			)
			//查询币种
			g.Model("currency").Where("main_chain", "eth").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			//查询当前地址
			g.Model("address").Where("main_chain", "eth").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)
			//处理金额
			amount, _ := decimal.NewFromString(value.Amount)
			tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
			convertAmount := amount.Mul(tenDecimal).BigInt()

			//查询当前归集地址的最大Nonce
			MaxNonce, _ := g.Model("recharge").Where("main_chain", "eth").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")
			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
				hashResult, nonce, _ = rpc.TransferEth(string(privateKey), convertAmount, ethMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
			} else {
				hashResult, nonce, _ = rpc.TransferEthToken(string(privateKey), convertAmount, ethMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
			}

			if hashResult != nil {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
			}
		}
	}
}

// hecoRecharge 火币链任务归集
func hecoRecharge() {
	//判断一下是否有配置手续费地址和归集地址
	hecoMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.hecoMergeAddress")
	hecoFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.hecoFeeAddress")
	hecoFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.hecoFeeAddressPrivateKey")
	//如果未配置出金地址，退出
	if hecoMergeAddress.ConfigValue == "" || hecoFeeAddress.ConfigValue == "" || hecoFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("归集地址或手续费私钥地址未配置，退出归集")
		return
	}
	hecoGasLimit, _ := service.SysConfig.GetConfigByKey("sys.hecoGasLimit")
	hecoGasPrice, _ := service.SysConfig.GetConfigByKey("sys.hecoGasPrice")

	bnbFeePrivateKey, _ := library.DecryptByAes(hecoFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "heco").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("heco_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	minFee := gconv.Int64(hecoGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(hecoGasLimit.ConfigValue)

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
			if count, _ := g.Model("fee_list").Where("main_chain", "heco").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "heco").Where("withdraw_address", hecoFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.HecoTransferHt(string(bnbFeePrivateKey), big.NewInt(last), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount := gconv.Float64(last) / math.Pow10(18)
					g.Model("fee_list").Data(g.Map{"main_chain": "heco", "coin_name": "HT", "withdraw_address": hecoFeeAddress.ConfigValue, "address": value.ToAddress,
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
			g.Model("currency").Where("main_chain", "heco").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			//查询当前地址
			g.Model("address").Where("main_chain", "heco").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)
			//处理金额
			amount, _ := decimal.NewFromString(value.Amount)
			tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
			convertAmount := amount.Mul(tenDecimal).BigInt()

			//查询当前归集地址的最大Nonce
			MaxNonce, _ := g.Model("recharge").Where("main_chain", "heco").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")
			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
				hashResult, nonce, _ = rpc.HecoTransferHt(string(privateKey), convertAmount, hecoMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
			} else {
				hashResult, nonce, _ = rpc.HecoTransferToken(string(privateKey), convertAmount, hecoMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
			}

			if hashResult != nil {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
			}
		}
	}
}

// tronRecharge 波场链任务归集
func tronRecharge() {
	//判断一下是否有配置手续费地址和归集地址
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

// nacRecharge Nac链任务归集
func nacRecharge() {
	fmt.Println("开始执行归集任务")

	//判断一下是否有配置手续费地址和归集地址
	nacMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.nacMergeAddress")
	nacFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.nacFeeAddress")
	//如果未配置出金地址，退出
	if nacMergeAddress.ConfigValue == "" || nacFeeAddress.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("归集地址或手续费私钥地址未配置，退出归集")
		return
	}
	var (
		list       []*model.Recharge
		hashResult string
		err        error
	)
	err = g.Model("recharge").Where("main_chain", "nac").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}

	tronGasLimit, _ := service.SysConfig.GetConfigByKey("sys.nacFee")
	minFee, _ := decimal.NewFromString(tronGasLimit.ConfigValue)
	//minFee, _ := rpc.NacGetGasFee()
	for _, value := range list {

		//查询当前地址余额
		balance, _ := rpc.GetBalance(value.ToAddress, "1")
		amount, _ := decimal.NewFromString(value.Amount)

		var last float64
		if value.ContractAddress == "1" {
			can, _ := balance.Sub(minFee).Sub(amount).Float64()
			if can < 0 {
				last, _ = balance.Add(minFee).Sub(amount).Float64()
			} else {
				last = 0
			}
			fmt.Printf(" last ==%v \n", last)
		} else {
			can, _ := balance.Sub(minFee).Float64()
			if can < 0 {
				last, _ = minFee.Sub(balance).Float64()
			} else {
				last = 0
			}
			fmt.Printf(" last ==%v \n", last)
		}

		//需要先转手续费
		if last > 0 {
			//先查询是否已经打过了手续费
			hashResult, err = rpc.NacWithdraw(nacFeeAddress.ConfigValue, value.ToAddress, "1", gconv.String(last), "")
			if err != nil {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf(" 手续费打入失败 ==%v \n", err.Error())
			}
			if hashResult != "" {
				g.Model("fee_list").Data(g.Map{"main_chain": "nac", "coin_name": "nac", "withdraw_address": nacFeeAddress.ConfigValue, "address": value.ToAddress,
					"amount": last, "hash": hashResult, "recharge_id": value.Id}).Insert()
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 5}).Update()
			}
		} else {
			hashResult, err = rpc.NacWithdraw(value.ToAddress, nacMergeAddress.ConfigValue, value.ContractAddress, value.Amount, "")
			if err != nil {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf(" 归集打入失败 ==%v \n", err.Error())
			}
			if hashResult != "" {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "imputation_hash": hashResult}).Update()
			}
		}
	}
}

// wemixRecharge 火币链任务归集
func wemixRecharge() {
	//判断一下是否有配置手续费地址和归集地址
	wemixMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.wemixMergeAddress")
	wemixFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.wemixFeeAddress")
	wemixFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.wemixFeeAddressPrivateKey")
	//如果未配置出金地址，退出
	if wemixMergeAddress.ConfigValue == "" || wemixFeeAddress.ConfigValue == "" || wemixFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("归集地址或手续费私钥地址未配置，退出归集")
		return
	}
	wemixGasLimit, _ := service.SysConfig.GetConfigByKey("sys.wemixGasLimit")
	wemixGasPrice, _ := service.SysConfig.GetConfigByKey("sys.wemixGasPrice")

	bnbFeePrivateKey, _ := library.DecryptByAes(wemixFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "wemix").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("wemix_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	minFee := gconv.Int64(wemixGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(wemixGasLimit.ConfigValue)

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
			if count, _ := g.Model("fee_list").Where("main_chain", "wemix").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "wemix").Where("withdraw_address", wemixFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.WemixTransferWemix(string(bnbFeePrivateKey), big.NewInt(last), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount := gconv.Float64(last) / math.Pow10(18)
					g.Model("fee_list").Data(g.Map{"main_chain": "wemix", "coin_name": "WEMIX", "withdraw_address": wemixFeeAddress.ConfigValue, "address": value.ToAddress,
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
			g.Model("currency").Where("main_chain", "wemix").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			//查询当前地址
			g.Model("address").Where("main_chain", "wemix").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)
			//处理金额
			amount, _ := decimal.NewFromString(value.Amount)
			tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
			convertAmount := amount.Mul(tenDecimal).BigInt()

			//查询当前归集地址的最大Nonce
			MaxNonce, _ := g.Model("recharge").Where("main_chain", "wemix").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")
			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
				hashResult, nonce, _ = rpc.WemixTransferWemix(string(privateKey), convertAmount, wemixMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
			} else {
				hashResult, nonce, _ = rpc.WemixTransferToken(string(privateKey), convertAmount, wemixMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
			}

			if hashResult != nil {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
			}
		}
	}
}
