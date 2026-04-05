package task

import (
	"context"
	"fmt"
	commonModel "gfast/app/common/model"
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

// calculateEstimatedGasFee 计算预估的gas费用，考虑缓冲和多重交易场景
func calculateEstimatedGasFee(mainChain string, transactionCount int) decimal.Decimal {
	var gasLimitConfig, gasPriceConfig *commonModel.SysConfig

	switch mainChain {
	case "bsc":
		gasLimitConfig, _ = service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
		gasPriceConfig, _ = service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	case "eth":
		gasLimitConfig, _ = service.SysConfig.GetConfigByKey("sys.ethGasLimit")
		gasPriceConfig, _ = service.SysConfig.GetConfigByKey("sys.ethGasPrice")
	case "heco":
		gasLimitConfig, _ = service.SysConfig.GetConfigByKey("sys.hecoGasLimit")
		gasPriceConfig, _ = service.SysConfig.GetConfigByKey("sys.hecoGasPrice")
	case "wemix":
		gasLimitConfig, _ = service.SysConfig.GetConfigByKey("sys.wemixGasLimit")
		gasPriceConfig, _ = service.SysConfig.GetConfigByKey("sys.wemixGasPrice")
	default:
		g.Log().Printf("不支持的主链: %s", mainChain)
		return decimal.Zero
	}

	if gasLimitConfig == nil || gasPriceConfig == nil {
		g.Log().Printf("获取%s链gas配置失败", mainChain)
		return decimal.Zero
	}

	// 基础gas费用
	baseGasInt := gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(gasLimitConfig.ConfigValue)

	// 考虑交易数量（授权+归集等多重交易）并增加30%安全缓冲
	estimatedGasInt := baseGasInt * int64(transactionCount) * 130 / 100

	return decimal.NewFromInt(estimatedGasInt)
}

// 归集任务
func rechargeTask() {
	if g.Cfg().GetBool("bsc.address_recharge") || g.Cfg().GetBool("bsc.contract_recharge") {
		// 先处理已授权待归集的记录
		processAuthorizedRecharge()
		// 再处理普通归集
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
	if g.Cfg().GetBool("solana.address_recharge") || g.Cfg().GetBool("solana.contract_recharge") {
		solanaRecharge()
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
	bnbFeePrivateKey, _ := library.DecryptByAes(bnbFeeAddressPrivateKey.ConfigValue)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "bsc").Where("status", 1).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询未归集任务失败，%v", err)
	}
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)

	// 使用新的辅助函数计算预估gas费用，考虑可能的授权+归集两次交易
	minFee := calculateEstimatedGasFee("bsc", 2)

	for _, value := range list {

		//查询当前地址余额
		balanceWei, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)
		balance := decimal.NewFromBigInt(balanceWei, 0)

		var last decimal.Decimal
		if value.ContractAddress == "0x0000000000000000000000000000000000000000" {
			amount, _ := decimal.NewFromString(value.Amount)
			amountWei := amount.Mul(decimal.NewFromInt(int64(math.Pow10(18))))

			if balance.Sub(minFee).Sub(amountWei).IsNegative() {
				last = minFee.Add(amountWei).Sub(balance)
			} else {
				last = decimal.Zero
			}
		} else {
			if balance.Sub(minFee).IsNegative() {
				last = minFee.Sub(balance)
			} else {
				last = decimal.Zero
			}
		}

		//需要先转手续费
		if last.IsPositive() {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "bsc").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "bsc").Where("withdraw_address", bnbFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.TransferBnb(string(bnbFeePrivateKey), last.BigInt(), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount, _ := last.Div(decimal.NewFromInt(int64(math.Pow10(18)))).Float64()
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
			if g.Cfg().GetString("bsc.common_recharge") != "" {
				// 特殊归集逻辑：通过调用合约方法进行归集
				commonRechargeContract := g.Cfg().GetString("bsc.common_recharge")

				if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
					// BNB直接转账
					hashResult, nonce, _ = rpc.TransferBnb(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
				} else {
					// ERC20代币需要先检查授权，再调用归集合约
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("开始处理代币归集，地址: %s, 代币: %s, 数量: %s", value.ToAddress, currency.ContractAddress, value.Amount)

					// 检查是否已经授权给归集合约
					allowance, err := rpc.CheckTokenAllowance(currency.ContractAddress, value.ToAddress, commonRechargeContract)
					if err != nil {
						g.Log().File("merge_recharge.{Y-m-d}.log").Printf("检查授权失败: %v", err)
						continue
					}

					// 如果授权额度不足，需要先授权
					if allowance.Cmp(convertAmount) < 0 {
						// 检查是否已经在authorize_address表中记录了授权操作
						authorizeCount, _ := g.Model("authorize_address").Where("main_chain", "bsc").
							Where("address", value.ToAddress).
							Where("contract_address", commonRechargeContract).
							Where("coin_address", currency.ContractAddress).Count()

						if authorizeCount == 0 {
							g.Log().File("merge_recharge.{Y-m-d}.log").Printf("地址 %s 未授权代币 %s 给归集合约，开始授权", value.ToAddress, currency.ContractAddress)

							// 授权无限量代币（使用最大值）
							maxAmount := new(big.Int)
							maxAmount.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10) // 2^256 - 1

							approveHash, approveNonce, err := rpc.ApproveToken(string(privateKey), currency.ContractAddress, commonRechargeContract, maxAmount, gconv.Uint64(MaxNonce))
							if err != nil {
								g.Log().File("merge_recharge.{Y-m-d}.log").Printf("授权失败: %v", err)
								continue
							}

							if approveHash != nil {
								// 记录授权信息到authorize_address表
								g.Model("authorize_address").Data(g.Map{
									"main_chain":       "bsc",
									"contract_address": commonRechargeContract,
									"address":          value.ToAddress,
									"coin_name":        currency.Name,
									"coin_decimals":    currency.Decimals,
									"coin_address":     currency.ContractAddress,
									"num":              "115792089237316195423570985008687907853269984665640564039457584007913129639935",
									"authorize_hash":   approveHash,
								}).Insert()

								// 更新充值记录状态为授权中
								g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 6, "nonce": approveNonce, "imputation_hash": approveHash}).Update()
								g.Log().File("merge_recharge.{Y-m-d}.log").Printf("授权成功，hash: %s", approveHash)
							}
						} else {
							g.Log().File("merge_recharge.{Y-m-d}.log").Printf("地址 %s 已有授权记录，等待确认中", value.ToAddress)
						}
					} else {
						// 已授权，直接调用归集合约
						g.Log().File("merge_recharge.{Y-m-d}.log").Printf("地址 %s 已授权，开始调用归集合约", value.ToAddress)

						rechargeHash, rechargeNonce, err := rpc.CallRechargeContract(string(privateKey), commonRechargeContract, convertAmount, gconv.Uint64(MaxNonce))
						if err != nil {
							g.Log().File("merge_recharge.{Y-m-d}.log").Printf("调用归集合约失败: %v", err)
							continue
						}

						if rechargeHash != nil {
							hashResult = rechargeHash
							nonce = rechargeNonce
							g.Log().File("merge_recharge.{Y-m-d}.log").Printf("归集合约调用成功，hash: %s", rechargeHash)
						}
					}
				}
			} else {
				if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
					hashResult, nonce, _ = rpc.TransferBnb(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
				} else {
					hashResult, nonce, _ = rpc.TransferToken(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
				}
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
	minFeeInt := gconv.Int64(ethGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(ethGasLimit.ConfigValue)
	minFee := decimal.NewFromInt(minFeeInt)

	for _, value := range list {

		//查询当前地址余额
		balanceWei, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)
		balance := decimal.NewFromBigInt(balanceWei, 0)

		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询0 %v 余额为 %v", value.ToAddress, balance)

		var last decimal.Decimal
		if value.ContractAddress == "0x0000000000000000000000000000000000000000" {
			amount, _ := decimal.NewFromString(value.Amount)
			amountWei := amount.Mul(decimal.NewFromInt(int64(math.Pow10(18))))

			if balance.Sub(minFee).Sub(amountWei).IsNegative() {
				last = minFee.Add(amountWei).Sub(balance)
			} else {
				last = decimal.Zero
			}
		} else {
			if balance.Sub(minFee).IsNegative() {
				last = minFee.Sub(balance)
			} else {
				last = decimal.Zero
			}
		}

		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询1 %v 余额为 %v", value.ToAddress, balance)

		//需要先转手续费
		if last.IsPositive() {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "eth").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "eth").Where("withdraw_address", ethFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询4 %v 余额为 %v", value.ToAddress, balance)
				hashResult, nonce, _ := rpc.TransferEth(string(ethFeePrivateKey), last.BigInt(), value.ToAddress, gconv.Uint64(MaxNonce))
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询65 %v 余额为 %v", value.ToAddress, balance)
				if hashResult != nil {
					amount, _ := last.Div(decimal.NewFromInt(int64(math.Pow10(18)))).Float64()
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
			if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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
	minFeeInt := gconv.Int64(hecoGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(hecoGasLimit.ConfigValue)
	minFee := decimal.NewFromInt(minFeeInt)

	for _, value := range list {

		//查询当前地址余额
		balanceWei, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)
		balance := decimal.NewFromBigInt(balanceWei, 0)

		var last decimal.Decimal
		if value.ContractAddress == "0x0000000000000000000000000000000000000000" {
			amount, _ := decimal.NewFromString(value.Amount)
			amountWei := amount.Mul(decimal.NewFromInt(int64(math.Pow10(18))))

			if balance.Sub(minFee).Sub(amountWei).IsNegative() {
				last = minFee.Add(amountWei).Sub(balance)
			} else {
				last = decimal.Zero
			}
		} else {
			if balance.Sub(minFee).IsNegative() {
				last = minFee.Sub(balance)
			} else {
				last = decimal.Zero
			}
		}

		//需要先转手续费
		if last.IsPositive() {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "heco").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "heco").Where("withdraw_address", hecoFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.HecoTransferHt(string(bnbFeePrivateKey), last.BigInt(), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount, _ := last.Div(decimal.NewFromInt(int64(math.Pow10(18)))).Float64()
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
			if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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
				if err != nil {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("TRON归集失败: %v", err)
					continue
				}
			} else {
				//处理金额
				amount, _ := decimal.NewFromString(value.Amount)
				tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
				convertAmount := amount.Mul(tenDecimal).BigInt()

				txId, err = tronClient.TransferContract(string(privateKey), value.ToAddress, tronMergeAddress.ConfigValue, value.ContractAddress, convertAmount, tronGasFee)
				if err != nil {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("TRON代币归集失败: %v", err)
					continue
				}
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
				last, _ = minFee.Add(amount).Sub(balance).Float64()
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
	minFeeInt := gconv.Int64(wemixGasPrice.ConfigValue) * gconv.Int64(math.Pow10(9)) * gconv.Int64(wemixGasLimit.ConfigValue)
	minFee := decimal.NewFromInt(minFeeInt)

	for _, value := range list {

		//查询当前地址余额
		balanceWei, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)
		balance := decimal.NewFromBigInt(balanceWei, 0)

		var last decimal.Decimal
		if value.ContractAddress == "0x0000000000000000000000000000000000000000" {
			amount, _ := decimal.NewFromString(value.Amount)
			amountWei := amount.Mul(decimal.NewFromInt(int64(math.Pow10(18))))

			if balance.Sub(minFee).Sub(amountWei).IsNegative() {
				last = minFee.Add(amountWei).Sub(balance)
			} else {
				last = decimal.Zero
			}
		} else {
			if balance.Sub(minFee).IsNegative() {
				last = minFee.Sub(balance)
			} else {
				last = decimal.Zero
			}
		}

		//需要先转手续费
		if last.IsPositive() {
			//先查询是否已经打过了手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "wemix").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				MaxNonce, _ := g.Model("fee_list").Where("main_chain", "wemix").Where("withdraw_address", wemixFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")
				hashResult, nonce, _ := rpc.WemixTransferWemix(string(bnbFeePrivateKey), last.BigInt(), value.ToAddress, gconv.Uint64(MaxNonce))

				if hashResult != nil {
					amount, _ := last.Div(decimal.NewFromInt(int64(math.Pow10(18)))).Float64()
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
			if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
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

// solanaRecharge Solana链任务归集
func solanaRecharge() {
	g.Log().File("merge_recharge.{Y-m-d}.log").Println("开始Solana归集")

	solanaMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.solanaMergeAddress")
	solanaFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.solanaFeeAddress")
	solanaFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.solanaFeeAddressPrivateKey")

	if solanaMergeAddress == nil || solanaMergeAddress.ConfigValue == "" ||
		solanaFeeAddress == nil || solanaFeeAddress.ConfigValue == "" ||
		solanaFeeAddressPrivateKey == nil || solanaFeeAddressPrivateKey.ConfigValue == "" {
		g.Log().File("merge_recharge.{Y-m-d}.log").Println("Solana归集地址或手续费私钥未配置，退出归集")
		return
	}

	solanaFeePrivateKey, _ := library.DecryptByAes(solanaFeeAddressPrivateKey.ConfigValue)

	// Solana单笔交易手续费约 5000 lamports (0.000005 SOL)
	// 为SPL Token转账预留更多（可能需要创建ATA: ~0.002 SOL）
	// 设定最小手续费为 0.003 SOL = 3_000_000 lamports
	minFeeLamports := int64(3_000_000)

	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "solana").Where("status", 1).Limit(20).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana查询未归集任务失败: %v", err)
		return
	}
	g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana待归集记录: %d 条", len(list))

	for _, value := range list {
		// 查询用户地址SOL余额
		balanceLamports, err := rpc.SolanaClient.GetBalance(value.ToAddress)
		if err != nil {
			g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana查询余额失败 addr=%v err=%v", value.ToAddress, err)
			continue
		}
		balance := int64(balanceLamports)

		var needFee int64
		if value.ContractAddress == rpc.SOLNativeMint {
			// SOL原生转账：需要手续费 + 转出金额
			amount, _ := decimal.NewFromString(value.Amount)
			amountLamports := amount.Mul(decimal.NewFromInt(rpc.LamportsPerSOL)).IntPart()
			if balance-minFeeLamports-amountLamports < 0 {
				needFee = minFeeLamports + amountLamports - balance
			}
		} else {
			// SPL Token转账：只需要SOL手续费
			if balance-minFeeLamports < 0 {
				needFee = minFeeLamports - balance
			}
		}

		if needFee > 0 {
			// 需要先转手续费
			if count, _ := g.Model("fee_list").Where("main_chain", "solana").Where("address", value.ToAddress).Where("status", 1).Count(); count == 0 {
				feeAmountSOL := decimal.NewFromInt(needFee).Div(decimal.NewFromInt(rpc.LamportsPerSOL))
				txSig, err := rpc.SolanaTransferSOL(string(solanaFeePrivateKey), value.ToAddress, feeAmountSOL)
				if err != nil {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana手续费转账失败 addr=%v err=%v", value.ToAddress, err)
					continue
				}
				if txSig != "" {
					feeAmount, _ := feeAmountSOL.Float64()
					g.Model("fee_list").Data(g.Map{
						"main_chain":       "solana",
						"coin_name":        "SOL",
						"withdraw_address": solanaFeeAddress.ConfigValue,
						"address":          value.ToAddress,
						"amount":           feeAmount,
						"hash":             txSig,
						"recharge_id":      value.Id,
					}).Insert()
					g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 5}).Update()
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana手续费已转出 addr=%v sig=%v amount=%v", value.ToAddress, txSig, feeAmount)
				}
			}
		} else {
			// 余额足够，直接归集
			var (
				address  *model.Address
				currency *model.Currency
			)
			g.Model("currency").Where("main_chain", "solana").Where("contract_address", value.ContractAddress).FindScan(&currency)
			if currency == nil {
				continue
			}
			g.Model("address").Where("main_chain", "solana").Where("address", value.ToAddress).FindScan(&address)
			if address == nil {
				continue
			}
			privateKey, _ := library.DecryptByAes(address.PrivateKey)

			var txSig string
			if value.ContractAddress == rpc.SOLNativeMint {
				// SOL归集：转出充值金额（扣除预留租金）
				amount, _ := decimal.NewFromString(value.Amount)
				txSig, err = rpc.SolanaTransferSOL(string(privateKey), solanaMergeAddress.ConfigValue, amount)
			} else {
				// SPL Token归集
				amount, _ := decimal.NewFromString(value.Amount)
				txSig, err = rpc.SolanaTransferSPLToken(string(privateKey), solanaMergeAddress.ConfigValue, value.ContractAddress, amount, currency.Decimals)
			}

			if err != nil {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana归集失败 id=%v err=%v", value.Id, err)
				continue
			}
			if txSig != "" {
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "imputation_hash": txSig}).Update()
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("Solana归集成功 id=%v sig=%v", value.Id, txSig)
			}
		}
	}
}

// processAuthorizedRecharge 处理已授权待归集的记录
func processAuthorizedRecharge() {
	// 查询状态为6(授权中)的记录
	var list []*model.Recharge
	err := g.Model("recharge").Where("main_chain", "bsc").Where("status", 6).Scan(&list)
	if err != nil {
		g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询授权中任务失败，%v", err)
		return
	}

	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	defer client.Close()

	commonRechargeContract := g.Cfg().GetString("bsc.common_recharge")
	bnbMergeAddress, _ := service.SysConfig.GetConfigByKey("sys.bnbMergeAddress")

	for _, value := range list {
		// 检查授权交易是否已确认
		if value.ImputationHash != "" {
			receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(value.ImputationHash))
			if err != nil {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("查询授权交易状态失败: %v", err)
				continue
			}

			// 如果交易成功确认
			if receipt.Status == 1 {
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("授权交易已确认，开始执行归集，地址: %s", value.ToAddress)

				var (
					currency *model.Currency
					address  *model.Address
				)

				// 查询币种和地址信息
				g.Model("currency").Where("main_chain", "bsc").Where("contract_address", value.ContractAddress).FindScan(&currency)
				if currency == nil {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("未找到币种信息，合约地址: %s", value.ContractAddress)
					continue
				}

				g.Model("address").Where("main_chain", "bsc").Where("address", value.ToAddress).FindScan(&address)
				if address == nil {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("未找到地址信息，地址: %s", value.ToAddress)
					continue
				}

				// 授权后重新检查余额，确认手续费是否足够进行归集
				balanceWei, _ := client.BalanceAt(context.Background(), common.HexToAddress(value.ToAddress), nil)
				currentBalance := decimal.NewFromBigInt(balanceWei, 0)

				// 使用辅助函数计算归集所需的预估gas费用（只需要1次归集交易）
				estimatedFee := calculateEstimatedGasFee("bsc", 1)

				// 检查余额是否足够支付归集交易
				if currentBalance.LessThan(estimatedFee) {
					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("授权后余额不足，当前余额: %s, 预估手续费: %s, 需要补充手续费，地址: %s",
						currentBalance.String(), estimatedFee.String(), value.ToAddress)

					// 计算需要补充的手续费
					additionalFee := estimatedFee.Sub(currentBalance)
					bnbFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.bnbFeeAddress")
					bnbFeeAddressPrivateKey, _ := service.SysConfig.GetConfigByKey("sys.bnbFeeAddressPrivateKey")
					bnbFeePrivateKey, _ := library.DecryptByAes(bnbFeeAddressPrivateKey.ConfigValue)

					// 获取手续费地址的最大nonce
					feeMaxNonce, _ := g.Model("fee_list").Where("main_chain", "bsc").Where("withdraw_address", bnbFeeAddress.ConfigValue).OrderDesc("id").Value("nonce")

					// 补充手续费
					hashResult, nonce, err := rpc.TransferBnb(string(bnbFeePrivateKey), additionalFee.BigInt(), value.ToAddress, gconv.Uint64(feeMaxNonce))
					if err != nil || hashResult == nil {
						g.Log().File("merge_recharge.{Y-m-d}.log").Printf("补充手续费失败: %v, 地址: %s", err, value.ToAddress)
						continue
					}

					// 记录补充的手续费
					additionalFeeEth, _ := additionalFee.Div(decimal.NewFromInt(int64(math.Pow10(18)))).Float64()
					g.Model("fee_list").Data(g.Map{
						"main_chain":       "bsc",
						"coin_name":        "bnb",
						"withdraw_address": bnbFeeAddress.ConfigValue,
						"address":          value.ToAddress,
						"amount":           additionalFeeEth,
						"hash":             hashResult,
						"nonce":            nonce,
						"recharge_id":      value.Id,
						"remark":           "授权后补充手续费",
					}).Insert()

					g.Log().File("merge_recharge.{Y-m-d}.log").Printf("成功补充手续费: %.6f BNB, hash: %s", additionalFeeEth, hashResult)

					// 暂停处理，等待下一轮任务确认手续费到账后再进行归集
					continue
				}

				privateKey, _ := library.DecryptByAes(address.PrivateKey)

				// 计算转账金额
				amount, _ := decimal.NewFromString(value.Amount)
				tenDecimal := decimal.NewFromInt(gconv.Int64(math.Pow(10, float64(currency.Decimals))))
				convertAmount := amount.Mul(tenDecimal).BigInt()

				// 查询当前地址的最大Nonce
				MaxNonce, _ := g.Model("recharge").Where("main_chain", "bsc").Where("to_address", value.ToAddress).WhereIn("status", [3]int{2, 3, 4}).OrderDesc("id").Value("nonce")

				// 调用归集合约
				if commonRechargeContract != "" {
					rechargeHash, rechargeNonce, err := rpc.CallRechargeContract(string(privateKey), commonRechargeContract, convertAmount, gconv.Uint64(MaxNonce))
					if err != nil {
						g.Log().File("merge_recharge.{Y-m-d}.log").Printf("调用归集合约失败: %v", err)
						// 将状态改为归集失败
						g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 4}).Update()
						continue
					}

					if rechargeHash != nil {
						// 更新为归集上链中状态
						g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": rechargeNonce, "imputation_hash": rechargeHash}).Update()
						g.Log().File("merge_recharge.{Y-m-d}.log").Printf("归集合约调用成功，hash: %s", rechargeHash)
					}
				} else {
					// 普通归集方式
					if currency.ContractAddress == "0x0000000000000000000000000000000000000000" {
						hashResult, nonce, _ := rpc.TransferBnb(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, gconv.Uint64(MaxNonce))
						if hashResult != nil {
							g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
						}
					} else {
						hashResult, nonce, _ := rpc.TransferToken(string(privateKey), convertAmount, bnbMergeAddress.ConfigValue, currency.ContractAddress, gconv.Uint64(MaxNonce))
						if hashResult != nil {
							g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 2, "nonce": nonce, "imputation_hash": hashResult}).Update()
						}
					}
				}
			} else if receipt.Status == 0 {
				// 授权交易失败，重置状态为待归集
				g.Log().File("merge_recharge.{Y-m-d}.log").Printf("授权交易失败，重置状态，地址: %s", value.ToAddress)
				g.Model("recharge").Where("id", value.Id).Data(g.Map{"status": 1}).Update()
				// 删除失败的授权记录
				g.Model("authorize_address").Where("main_chain", "bsc").Where("address", value.ToAddress).Where("authorize_hash", value.ImputationHash).Delete()
			}
		}
	}
}
