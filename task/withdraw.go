package task

import (
	"fmt"
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

const solanaWithdrawBatchSizePerCurrency = 20

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
	if g.Config().GetBool("solana.address_recharge") || g.Config().GetBool("solana.contract_recharge") {
		solanaWithdraw()
	}
}

func bscWithdraw() {
	g.Log().File("withdraw.{Y-m-d}.log").Printf("开始执行提现任务")

	bnbWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbWithdrawAddressPrivateKey")
	//如果未配置出金地址，退出
	if bnbWithdrawPrivateKeyConfig.ConfigValue == "" {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("未配置出金地址私钥")
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
				if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
					hashResult, nonce, _ = rpc.TransferBnb(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce))
				} else if common.HexToAddress(currency.ContractAddress) == common.HexToAddress("0x70b8ce59baE2FdB419C9489813ee51D14028b8d9") {
					hashResult, nonce, _ = rpc.TransferTokenOr(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
				} else {
					if value.Function != "" && value.FunctionAddress != "" {
						hashResult, nonce, _ = rpc.SpecifyTransferToken(string(bnbPrivateKey), convertAmount, value.Address, gconv.Uint64(MaxNonce), value.Function, value.FunctionAddress)
					} else {
						hashResult, nonce, _ = rpc.TransferToken(string(bnbPrivateKey), convertAmount, value.Address, currency.ContractAddress, gconv.Uint64(MaxNonce))
					}
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
		if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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
//			if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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

func solanaAmountToBaseUnits(amount decimal.Decimal, decimals int) (uint64, error) {
	if decimals < 0 || decimals > 255 {
		return 0, fmt.Errorf("invalid Solana currency decimals: %d", decimals)
	}
	scaled := amount.Mul(decimal.New(1, int32(decimals))).Truncate(0)
	if !scaled.IsPositive() {
		return 0, fmt.Errorf("Solana transfer amount is less than one base unit: amount=%s decimals=%d", amount.String(), decimals)
	}
	integer := scaled.BigInt()
	if integer.Sign() < 0 || integer.BitLen() > 64 {
		return 0, fmt.Errorf("Solana transfer amount exceeds uint64: amount=%s decimals=%d", amount.String(), decimals)
	}
	return integer.Uint64(), nil
}

func solanaBaseUnitsToAmount(amount uint64, decimals int) string {
	return decimal.NewFromBigInt(new(big.Int).SetUint64(amount), int32(-decimals)).String()
}

func buildSolanaWithdrawTransfers(value *model.Withdraw, currency *model.Currency) ([]rpc.SolanaTransfer, string, string, error) {
	if value == nil || currency == nil {
		return nil, "", "", fmt.Errorf("Solana withdrawal or currency is nil")
	}
	totalUnits, err := solanaAmountToBaseUnits(decimal.NewFromFloat(value.Amount), currency.Decimals)
	if err != nil {
		return nil, "", "", err
	}
	if currency.WithdrawSplitEnabled == 0 {
		return []rpc.SolanaTransfer{{ToAddress: value.Address, Amount: totalUnits}}, "", "0", nil
	}
	if currency.WithdrawSplitEnabled != 1 {
		return nil, "", "", fmt.Errorf("invalid Solana withdraw split switch: currency=%s value=%d", currency.Name, currency.WithdrawSplitEnabled)
	}
	if currency.WithdrawSplitAddress == "" {
		return nil, "", "", fmt.Errorf("Solana withdraw split address is empty: currency=%s", currency.Name)
	}
	if currency.WithdrawSplitAddress == value.Address {
		return nil, "", "", fmt.Errorf("Solana withdraw split address equals user address: currency=%s address=%s", currency.Name, value.Address)
	}
	fixedAmountText := currency.WithdrawSplitAmount
	if fixedAmountText == "" {
		fixedAmountText = "0"
	}
	fixedSplitAmount, err := decimal.NewFromString(fixedAmountText)
	if err != nil || fixedSplitAmount.IsNegative() {
		return nil, "", "", fmt.Errorf("invalid Solana withdraw split amount: currency=%s amount=%s", currency.Name, currency.WithdrawSplitAmount)
	}
	if currency.WithdrawSplitBps >= 10000 {
		return nil, "", "", fmt.Errorf("invalid Solana withdraw split bps: currency=%s bps=%d", currency.Name, currency.WithdrawSplitBps)
	}
	hasFixedAmount := fixedSplitAmount.IsPositive()
	hasRate := currency.WithdrawSplitBps > 0
	if hasFixedAmount == hasRate {
		return nil, "", "", fmt.Errorf("Solana fixed split amount and split bps must be mutually exclusive: currency=%s", currency.Name)
	}

	var splitUnits uint64
	if hasRate {
		numerator := new(big.Int).Mul(new(big.Int).SetUint64(totalUnits), big.NewInt(int64(currency.WithdrawSplitBps)))
		splitInteger := new(big.Int).Div(numerator, big.NewInt(10000))
		if splitInteger.Sign() <= 0 {
			return nil, "", "", fmt.Errorf("Solana percentage split amount is less than one base unit: currency=%s total=%d bps=%d", currency.Name, totalUnits, currency.WithdrawSplitBps)
		}
		splitUnits = splitInteger.Uint64()
	} else {
		scaledSplitAmount := fixedSplitAmount.Mul(decimal.New(1, int32(currency.Decimals)))
		if !scaledSplitAmount.Equal(scaledSplitAmount.Truncate(0)) {
			return nil, "", "", fmt.Errorf("Solana withdraw split amount exceeds currency precision: currency=%s amount=%s decimals=%d", currency.Name, fixedSplitAmount.String(), currency.Decimals)
		}
		splitUnits, err = solanaAmountToBaseUnits(fixedSplitAmount, currency.Decimals)
		if err != nil {
			return nil, "", "", fmt.Errorf("invalid Solana withdraw split amount: currency=%s: %v", currency.Name, err)
		}
	}
	if splitUnits >= totalUnits {
		return nil, "", "", fmt.Errorf("Solana withdraw split amount must be less than total: currency=%s total=%d split=%d", currency.Name, totalUnits, splitUnits)
	}
	actualSplitAmount := solanaBaseUnitsToAmount(splitUnits, currency.Decimals)
	return []rpc.SolanaTransfer{
		{ToAddress: value.Address, Amount: totalUnits - splitUnits},
		{ToAddress: currency.WithdrawSplitAddress, Amount: splitUnits},
	}, currency.WithdrawSplitAddress, actualSplitAmount, nil
}

func solanaWithdraw() {
	g.Log().File("withdraw.{Y-m-d}.log").Printf("开始处理Solana出金任务")
	solanaWithdrawPrivateKeyConfig, _ := service.SysConfig.GetConfigByKey("sys.solanaWithdrawAddressPrivateKey")
	solanaWithdrawAddressConfig, _ := service.SysConfig.GetConfigByKey("sys.solanaWithdrawAddress")
	if solanaWithdrawPrivateKeyConfig == nil || solanaWithdrawPrivateKeyConfig.ConfigValue == "" ||
		solanaWithdrawAddressConfig == nil || solanaWithdrawAddressConfig.ConfigValue == "" {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana未配置出金地址或私钥")
		return
	}

	// Each currency gets its own batch. A currency with insufficient balance remains
	// pending for retry, but must not occupy the processing window of other currencies.
	contractAddresses, err := g.Model("withdraw").
		Where("main_chain", "solana").
		Where("status", 2).
		Group("contract_address").
		Array("contract_address")
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("查询Solana待出金币种失败 err=%v", err)
		return
	}
	if len(contractAddresses) == 0 {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana未找到需要出金的记录")
		return
	}

	solanaPrivateKey, _ := library.DecryptByAes(solanaWithdrawPrivateKeyConfig.ConfigValue)
	solanaAddress := solanaWithdrawAddressConfig.ConfigValue

	for _, contractAddress := range contractAddresses {
		ids, queryErr := g.Model("withdraw").
			Where("main_chain", "solana").
			Where("status", 2).
			Where("contract_address", contractAddress.String()).
			Order("id ASC").
			Limit(solanaWithdrawBatchSizePerCurrency).
			Array("id")
		if queryErr != nil {
			g.Log().File("withdraw.{Y-m-d}.log").Printf("查询Solana待出金记录失败 contract_address=%v err=%v", contractAddress.String(), queryErr)
			continue
		}

		for _, id := range ids {
			var value *model.Withdraw
			var currency *model.Currency
			g.Model("withdraw").Where("id", id).Where("status", 2).FindScan(&value)
			if value == nil {
				continue
			}
			g.Model("currency").Where("main_chain", "solana").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}

			hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, value.Status, value.Nonce1)
			if hashKey != value.HashKey {
				g.Model("withdraw").Data(g.Map{"status": 0}).Where("id", value.Id).Update()
				continue
			}

			transfers, splitAddress, splitAmount, buildErr := buildSolanaWithdrawTransfers(value, currency)
			if buildErr != nil {
				g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现分账配置无效 id=%v err=%v", id, buildErr)
				continue
			}

			var prepared *rpc.PreparedSolanaTransaction
			if value.ContractAddress == rpc.SOLNativeMint {
				prepared, err = rpc.SolanaClient.PrepareSOLTransfers(string(solanaPrivateKey), transfers)
			} else {
				prepared, err = rpc.SolanaClient.PrepareSPLTokenTransfers(string(solanaPrivateKey), value.ContractAddress, transfers)
			}

			if err != nil || prepared == nil {
				g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现失败 id=%v err=%v", id, err)
				continue
			}

			// 先持久化本地已确定的 signature，再广播。条件更新保证多实例只会有一个执行者发送。
			hashKey = library.Md5Data(value.Address, value.ContractAddress, value.Amount, 3, value.Nonce1)
			result, updateErr := g.Model("withdraw").Data(g.Map{
				"withdraw_address": solanaAddress,
				"hashKey":          hashKey,
				"hash":             prepared.Signature,
				"nonce":            prepared.LastValidBlockHeight,
				"split_address":    splitAddress,
				"split_amount":     splitAmount,
				"status":           3,
			}).Where("id", value.Id).Where("status", 2).Update()
			if updateErr != nil {
				g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现保存签名失败 id=%v signature=%v err=%v", id, prepared.Signature, updateErr)
				continue
			}
			affected, affectedErr := result.RowsAffected()
			if affectedErr != nil || affected != 1 {
				continue
			}

			if _, err = rpc.SolanaClient.SendPreparedTransaction(prepared); err != nil {
				if rpc.IsSolanaRPCError(err) {
					// RPC 已明确拒绝（例如预执行失败），交易没有被接收，可安全标记失败。
					failedHashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, 4, value.Nonce1)
					if _, updateErr = g.Model("withdraw").Data(g.Map{
						"hashKey": failedHashKey,
						"status":  4,
					}).Where("id", value.Id).Where("status", 3).Where("hash", prepared.Signature).Update(); updateErr != nil {
						g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现广播被拒绝且回写失败 id=%v signature=%v err=%v updateErr=%v", id, prepared.Signature, err, updateErr)
						continue
					}
					g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现广播被RPC拒绝 id=%v signature=%v err=%v", id, prepared.Signature, err)
					continue
				}

				// 网络超时或响应丢失时结果不确定，保持上链中并按已保存的 signature 核对。
				g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana提现广播结果不确定 id=%v signature=%v err=%v", id, prepared.Signature, err)
			}
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
			if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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
