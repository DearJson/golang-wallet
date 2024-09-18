package task

import (
	"fmt"
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// 出金任务
func withdrawTask() {
	if g.Config().GetBool("bsc.contract_recharge") || g.Config().GetBool("bsc.address_recharge") {
		bscWithdraw()
	}
	if g.Config().GetBool("tron.contract_recharge") || g.Config().GetBool("tron.address_recharge") {
		tronWithdraw()
	}
	//if g.Config().GetBool("heco.contract_recharge") || g.Config().GetBool("heco.address_recharge") {
	//	hecoWithdraw()
	//}
	if g.Config().GetBool("wemix.contract_recharge") || g.Config().GetBool("wemix.address_recharge") {
		wemixWithdraw()
	}
	if g.Config().GetBool("eth.contract_recharge") || g.Config().GetBool("eth.address_recharge") {
		ethWithdraw()
	}
	if g.Cfg().GetBool("nac.address_recharge") {
		nacWithdraw()
	}
}

func bscWithdraw() {
	bnbWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbWithdrawAddressPrivateKey")
	//如果未配置出金地址，退出
	if bnbWithdrawPrivateKeyConfig.ConfigValue == "" {
		return
	}

	//查询所有待出金的任务
	ids, err := g.Model("withdraw").Where("main_chain", "bsc").Where("status", 2).Limit(20).Array("id")
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}

	bnbPrivateKey, _ := library.DecryptByAes(bnbWithdrawPrivateKeyConfig.ConfigValue)
	for _, id := range ids {
		var value *model.Withdraw
		var currency *model.Currency
		var hashResult interface{}
		var nonce uint64
		//查询状态
		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
		if value == nil {
			continue
		}
		//查询币种
		g.Model("currency").Where("main_chain", "bsc").Where("contract_address", value.ContractAddress).FindScan(&currency)
		if currency == nil {
			continue
		}

		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
		//如果验证加密key不一致，直接该状态为0
		if hashKey != value.HashKey {
			g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
			continue
		}

		//处理金额
		amount := decimal.NewFromFloat(value.Amount)
		tenDecimal := decimal.NewFromFloat(math.Pow(10, float64(currency.Decimals)))
		convertAmount := amount.Mul(tenDecimal).BigInt()
		withdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.bnbWithdrawAddress")

		MaxNonce, _ := g.Model("withdraw").Where("main_chain", "bsc").Where("withdraw_address", withdrawAddress.ConfigValue).WhereIn("status", [3]int{3, 4, 5}).Max("nonce")

		if currency.Decimals == 0 {
			//这里是提现NFT卡牌
			hashResult, nonce, _ = rpc.SafeTransferFrom(string(bnbPrivateKey), value.Address, currency.ContractAddress, big.NewInt(value.TokenId), big.NewInt(int64(value.Amount)), gconv.Uint64(MaxNonce))
		} else {
			//交换代币
			if value.SwapRoute != "" {
				hashResult, nonce, _ = rpc.SwapToken(string(bnbPrivateKey), convertAmount, value.Address, value.SwapRoute, value.SwapPath, gconv.Uint64(MaxNonce))
			} else {
				if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
					hashResult, nonce, _ = rpc.TransferBnb(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce))
				} else if common.HexToAddress(currency.ContractAddress) == common.HexToAddress("0x70b8ce59baE2FdB419C9489813ee51D14028b8d9") {
					hashResult, nonce, _ = rpc.TransferTokenOr(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
				} else {
					hashResult, nonce, _ = rpc.TransferToken(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
				}
			}
		}

		if hashResult != nil {
			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
			g.Model("withdraw").Data(g.Map{"withdraw_address": withdrawAddress.ConfigValue, "hashKey": hashKey, "hash": hashResult, "nonce": nonce, "status": 3}).Where("id", value.Id).Update()
		}
	}
}

func ethWithdraw() {
	ethWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.ethWithdrawAddressPrivateKey")
	//如果未配置出金地址，退出
	if ethWithdrawPrivateKeyConfig.ConfigValue == "" {
		return
	}

	//查询所有待出金的任务
	ids, err := g.Model("withdraw").Where("main_chain", "eth").Where("status", 2).Limit(20).Array("id")
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}

	bnbPrivateKey, _ := library.DecryptByAes(ethWithdrawPrivateKeyConfig.ConfigValue)
	for _, id := range ids {
		var value *model.Withdraw
		var currency *model.Currency
		var hashResult interface{}
		var nonce uint64
		//查询状态
		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
		if value == nil {
			continue
		}
		//查询币种
		g.Model("currency").Where("main_chain", "eth").Where("contract_address", value.ContractAddress).FindScan(&currency)
		if currency == nil {
			continue
		}

		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
		//如果验证加密key不一致，直接该状态为0
		if hashKey != value.HashKey {
			g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
			continue
		}

		//处理金额
		amount := decimal.NewFromFloat(value.Amount)
		tenDecimal := decimal.NewFromFloat(math.Pow(10, float64(currency.Decimals)))
		convertAmount := amount.Mul(tenDecimal).BigInt()
		withdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.ethWithdrawAddress")

		MaxNonce, _ := g.Model("withdraw").Where("main_chain", "eth").Where("withdraw_address", withdrawAddress.ConfigValue).WhereIn("status", [3]int{3, 4, 5}).Max("nonce")
		//交换代币
		if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
			hashResult, nonce, _ = rpc.TransferEth(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce))
		} else {
			hashResult, nonce, _ = rpc.TransferEthToken(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
		}

		if hashResult != nil {
			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
			g.Model("withdraw").Data(g.Map{"withdraw_address": withdrawAddress.ConfigValue, "hashKey": hashKey, "hash": hashResult, "nonce": nonce, "status": 3}).Where("id", value.Id).Update()
		}
	}
}

//func hecoWithdraw() {
//	hecoWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.hecoWithdrawAddressPrivateKey")
//	//如果未配置出金地址，退出
//	if hecoWithdrawPrivateKeyConfig.ConfigValue == "" {
//		return
//	}
//
//	//查询所有待出金的任务
//	ids, err := g.Model("withdraw").Where("main_chain", "heco").Where("status", 2).Limit(20).Array("id")
//	if err != nil {
//		return
//	}
//	if len(ids) == 0 {
//		return
//	}
//
//	bnbPrivateKey, _ := library.DecryptByAes(hecoWithdrawPrivateKeyConfig.ConfigValue)
//	for _, id := range ids {
//		var value *model.Withdraw
//		var currency *model.Currency
//		var hashResult interface{}
//		var nonce uint64
//		//查询状态
//		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
//		if value == nil {
//			continue
//		}
//		//查询币种
//		g.Model("currency").Where("main_chain", "heco").Where("contract_address", value.ContractAddress).FindScan(&currency)
//		if currency == nil {
//			continue
//		}
//
//		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
//		//如果验证加密key不一致，直接该状态为0
//		if hashKey != value.HashKey {
//			g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
//		}
//
//		//处理金额
//		amount := decimal.NewFromFloat(value.Amount)
//		tenDecimal := decimal.NewFromFloat(math.Pow(10, float64(currency.Decimals)))
//		convertAmount := amount.Mul(tenDecimal).BigInt()
//		withdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.hecoWithdrawAddress")
//		MaxNonce, _ := g.Model("withdraw").Where("main_chain", "heco").Where("withdraw_address", withdrawAddress.ConfigValue).WhereIn("status", [3]int{3, 4, 5}).Max("nonce")
//
//		if currency.Decimals == 0 {
//			//这里是提现NFT卡牌
//			hashResult, nonce, _ = rpc.HecoMintNft(string(bnbPrivateKey), value.Address, currency.ContractAddress, big.NewInt(value.TokenId), value.Url, gconv.Uint64(MaxNonce))
//			fmt.Printf("%v,,,,%v \n", hashResult, err)
//		} else {
//			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
//				hashResult, nonce, _ = rpc.HecoTransferHt(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce))
//			} else {
//				hashResult, nonce, _ = rpc.HecoTransferToken(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
//			}
//		}
//
//		if hashResult != nil {
//			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
//			g.Model("withdraw").Data(g.Map{"withdraw_address": withdrawAddress.ConfigValue, "hashKey": hashKey, "hash": hashResult, "nonce": nonce, "status": 3}).Where("id", value.Id).Update()
//		}
//	}
//}

func tronWithdraw() {
	g.Log().File("withdraw.{Y-m-d}.log").Printf("开始处理波场出金任务")
	tronWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.tronWithdrawAddressPrivateKey")
	tronWithdrawAddressConfig, _ := service.SysConfig.GetConfigByKey("sys.tronWithdrawAddress")
	//如果未配置出金地址，退出
	if tronWithdrawPrivateKeyConfig.ConfigValue == "" || tronWithdrawAddressConfig.ConfigValue == "" {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("波场未配置出金地址或私钥")
		return
	}
	//查询所有待出金的任务
	ids, err := g.Model("withdraw").Where("main_chain", "tron").Where("status", 2).Limit(100).Array("id")
	if err != nil {
		return
	}
	if len(ids) == 0 {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("波场未找到需要出金的记录")
		return
	}

	var cache = service.Cache.New()
	tronClient, err := rpc.NewClient(gconv.String(cache.Get("tron_rpc_url")))
	if err != nil {
		g.Log().Printf("波场节点链接失败！！！%v", err)
		return
	}
	tronPrivateKey, _ := library.DecryptByAes(tronWithdrawPrivateKeyConfig.ConfigValue)
	tronAddress := tronWithdrawAddressConfig.ConfigValue
	tronGasLimit, _ := service.SysConfig.GetConfigByKey("sys.tronFee")
	tronGasFee := gconv.Int64(tronGasLimit.ConfigValue) * 1000000
	var (
		value    *model.Withdraw
		currency *model.Currency
		txId     string
	)
	for _, id := range ids {

		//查询状态
		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
		if value == nil {
			continue
		}
		//查询币种
		g.Model("currency").Where("main_chain", "tron").Where("contract_address", value.ContractAddress).FindScan(&currency)
		if currency == nil {
			continue
		}

		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
		//如果验证加密key不一致，直接该状态为0
		if hashKey != value.HashKey {
			g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
			continue
		}

		if currency.Decimals == 0 {

		} else {
			if currency.ContractAddress == "TBRop8PopYu8atWWez3g3ueVtSpseW78b6" {
				txId, err = tronClient.TransferTrx(string(tronPrivateKey), tronAddress, value.Address, decimal.NewFromFloat(value.Amount), value.TrxRemark)
			} else {
				//处理金额
				amount := big.NewFloat(value.Amount)
				tenDecimal := big.NewFloat(math.Pow(10, float64(currency.Decimals)))
				convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
				txId, err = tronClient.TransferContract(string(tronPrivateKey), tronAddress, value.Address, value.ContractAddress, convertAmount, tronGasFee)
			}
		}

		if err == nil {
			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
			g.Model("withdraw").Data(g.Map{"withdraw_address": tronAddress, "hashKey": hashKey, "hash": txId[2:], "status": 3}).Where("id", value.Id).Update()
		}
	}
}

func nacWithdraw() {
	nacWithdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.nacWithdrawAddress")
	//如果未配置出金地址，退出
	if nacWithdrawAddress.ConfigValue == "" {
		return
	}

	//查询所有待出金的任务
	ids, err := g.Model("withdraw").Where("main_chain", "nac").Where("status", 2).Limit(20).Array("id")
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}

	for _, id := range ids {
		var value *model.Withdraw
		//var nonce uint64
		//查询状态
		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
		if value == nil {
			continue
		}

		hashResult, err := rpc.NacWithdraw(nacWithdrawAddress.ConfigValue, value.Address, value.ContractAddress, gconv.String(value.Amount), "")
		if err != nil {
			fmt.Printf("操作是失败%v", err.Error())
		}

		if hashResult != "" {
			g.Model("withdraw").Data(g.Map{"withdraw_address": nacWithdrawAddress.ConfigValue, "hash": hashResult, "status": 3}).Where("id", value.Id).Update()
		}
	}
}

func wemixWithdraw() {
	wemixWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.wemixWithdrawAddressPrivateKey")
	//如果未配置出金地址，退出
	if wemixWithdrawPrivateKeyConfig.ConfigValue == "" {
		return
	}

	//查询所有待出金的任务
	ids, err := g.Model("withdraw").Where("main_chain", "wemix").Where("status", 2).Limit(20).Array("id")
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}

	bnbPrivateKey, _ := library.DecryptByAes(wemixWithdrawPrivateKeyConfig.ConfigValue)
	for _, id := range ids {
		var value *model.Withdraw
		var currency *model.Currency
		var hashResult interface{}
		var nonce uint64
		//查询状态
		g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
		if value == nil {
			continue
		}
		//查询币种
		g.Model("currency").Where("main_chain", "wemix").Where("contract_address", value.ContractAddress).FindScan(&currency)
		if currency == nil {
			continue
		}

		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
		//如果验证加密key不一致，直接该状态为0
		if hashKey != value.HashKey {
			g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
		}

		//处理金额
		amount := decimal.NewFromFloat(value.Amount)
		tenDecimal := decimal.NewFromFloat(math.Pow(10, float64(currency.Decimals)))
		convertAmount := amount.Mul(tenDecimal).BigInt()
		withdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.wemixWithdrawAddress")
		MaxNonce, _ := g.Model("withdraw").Where("main_chain", "wemix").Where("withdraw_address", withdrawAddress.ConfigValue).WhereIn("status", [3]int{3, 4, 5}).Max("nonce")

		if currency.Decimals == 0 {

		} else {
			if currency.ContractAddress == "0x1000000000000000000000000000000000000000" {
				hashResult, nonce, _ = rpc.WemixTransferWemix(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce))
			} else {
				hashResult, nonce, _ = rpc.WemixTransferToken(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
			}
		}

		if hashResult != nil {
			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
			g.Model("withdraw").Data(g.Map{"withdraw_address": withdrawAddress.ConfigValue, "hashKey": hashKey, "hash": hashResult, "nonce": nonce, "status": 3}).Where("id", value.Id).Update()
		}
	}
}
