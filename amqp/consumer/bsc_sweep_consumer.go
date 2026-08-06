package consumer

import (
	"encoding/hex"
	"encoding/json"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	sservice "gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"math"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
)

type BscSweepConsumer struct{}

// Consumer 实现币安扫块消费者
func (b *BscSweepConsumer) Consumer(dataByte []byte, key uint64) error {
	transfer := rpc.BscTransactions{}
	err := json.Unmarshal(dataByte, &transfer)
	if err != nil {
		return err
	} //json解析到结构体里面
	g.Log().File("consumer.{Y-m-d}.log").Printf("消费 %v", transfer.Hash)
	fromAddress := transfer.From
	//手续费地址
	bnbFeeAddress, err := service.SysConfig.GetConfigByKey("sys.bnbFeeAddress")
	if err != nil {
		return err
	}
	//提币地址
	bnbWithdrawAddress, err := service.SysConfig.GetConfigByKey("sys.bnbWithdrawAddress")
	if err != nil {
		return err
	}

	//判断一下这是合约充值还是获取地址充值
	if g.Config().GetBool("bsc.contract_recharge") && strings.EqualFold(transfer.To, g.Config().GetString("bsc.contract_address")) {
		err = contractRechargeHandle(&transfer)
	} else {
		if g.Config().GetBool("bsc.address_recharge") {
			if fromAddress == strings.ToLower(bnbWithdrawAddress.ConfigValue) || fromAddress == strings.ToLower(bnbFeeAddress.ConfigValue) {
				return nil
			}
			err = addressRechargeHandle(&transfer)
		}
	}
	return err
}

func contractRechargeHandle(transfer *rpc.BscTransactions) (err error) {
	//检查交易状态
	client := rpc.BscClient{}
	data, err := client.Init().GetTransferStatus(transfer.Hash)
	if err != nil {
		return err
	}
	bscStruct := rpc.BscTransferDetail{}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		g.Log().Printf("查询交易%v失败，%v# \n", transfer.Hash, err)
		return err
	}
	if bscStruct.Status == "0x0" {
		g.Log().Printf("交易状态是失败的，%v#不处理  %v#\n", transfer.Hash, err)
		return nil
	}
	fromAddress := transfer.From
	ctx := gctx.New()
	coinAddress, err := sservice.Currency.GetBnbCoinAddress(ctx)
	if err != nil {
		return err
	}

	functionName := library.SubStr(transfer.Input, 0, 10)
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)
	var (
		amount           string
		coinToken        string
		remarks          string
		rechargeType     int8
		contractAddress  string
		amount1          string
		coinToken1       string
		contractAddress1 string
		status           int8
		tokenId          string

		customeUser   string
		customeAmount string
		customeCoin   string
	)

	bhaRecharge, err := parseBHARechargeDepositEvent(bscStruct.Logs, transfer.To, transfer.From, coinAddress)
	if err != nil {
		return err
	}
	if bhaRecharge != nil {
		amount = bhaRecharge.Amount
		coinToken = bhaRecharge.CoinToken
		remarks = bhaRecharge.Remarks
		rechargeType = bhaRecharge.RechargeType
		contractAddress = bhaRecharge.ContractAddress
		amount1 = bhaRecharge.Amount1
		coinToken1 = bhaRecharge.CoinToken1
		contractAddress1 = bhaRecharge.ContractAddress1
		customeUser = bhaRecharge.CustomeUser
		customeAmount = bhaRecharge.CustomeAmount
		customeCoin = bhaRecharge.CustomeCoin
	} else if functionName == "0x68ca0399" {
		//单币充值
		contractAddress = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, 10, 64), "0"), 40, "0")
		inCoinAddress := false
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 74, 64), "0")
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()

		coinToken = coinAddress[contractAddress].Name
		fsd, err := hex.DecodeString(transfer.Input[266:])
		if err != nil {
			return err
		}
		remarks = gconv.String(fsd)
		rechargeType = 1
	} else if functionName == "0x166f95ee" {
		//备注是纯数字的 rechargeOne 单币充值
		contractAddress = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, 10, 64), "0"), 40, "0")
		inCoinAddress := false
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 74, 64), "0")
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		coinToken = coinAddress[contractAddress].Name

		remark11 := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 138, 64), "0")
		remarks = decimal.NewFromBigInt(library.HexToBigInt(remark11), 0).String()
		if err != nil {
			return err
		}
		rechargeType = 1
	} else if functionName == "0xf5125fa4" {
		//充值NFT
		transfer.Input = transfer.Input[2:]
		reader, err := os.Open("./abi/new_recharge/new_recharge.abi")
		if err != nil {
			return err
		}
		defer reader.Close()
		abi1, err := abi.JSON(reader)
		if err != nil {
			return err
		}
		if method, ok := abi1.Methods["rechargeCard"]; ok {
			decodeData, err := hex.DecodeString(transfer.Input)
			if err != nil {
				return err
			}
			params := map[string]interface{}{
				"_tokenAddress": nil,
				"_tokenId":      nil,
				"_num":          nil,
				"remarks":       nil,
			}
			err = method.Inputs.UnpackIntoMap(params, decodeData[4:])
			if err != nil {
				return err
			}
			contractAddress = strings.ToLower(gconv.String(params["_tokenAddress"]))
			coinToken = "NFT"
			amount = gconv.String(params["_num"])
			remarks = gconv.String(params["remarks"])
			tokenId = gconv.String(params["_tokenId"])
			rechargeType = 1
		} else {
			return nil
		}
	} else if functionName == "0xa6dba280" {
		//双币充值
		contractAddress = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, 10, 64), "0"), 40, "0")
		contractAddress1 = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, 138, 64), "0"), 40, "0")
		in1CoinAddress := false
		in2CoinAddress := false
		if _, ok := coinAddress[contractAddress]; ok {
			in1CoinAddress = true
		}
		if _, ok := coinAddress[contractAddress1]; ok {
			in2CoinAddress = true
		}
		if in1CoinAddress == false || in2CoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 74, 64), "0")
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		amount1String := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 64*3+10, 64), "0")
		amount1 = decimal.NewFromBigInt(library.HexToBigInt(amount1String), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress1].Decimals))), 18).String()
		rechargeType = 2
		coinToken = coinAddress[contractAddress].Name
		coinToken1 = coinAddress[contractAddress1].Name
		fsd, err := hex.DecodeString(transfer.Input[394:])
		if err != nil {
			return err
		}
		remarks = gconv.String(fsd)
	} else if functionName == "0x2e886a1d" {
		//通用充值方法
		reader, err := os.Open("./abi/new_recharge/new_recharge.abi")
		defer reader.Close()
		if err != nil {
			return err
		}
		tokenAbi, err := abi.JSON(reader)
		if err != nil {
			return err
		}
		decodedSig, _ := hex.DecodeString(transfer.Input[2:10])
		method, _ := tokenAbi.MethodById(decodedSig)
		decodedData, _ := hex.DecodeString(transfer.Input[10:])
		var functionInput rpc.RechargeFunctionInputs
		inputMap := make(map[string]interface{}, 0)
		_ = method.Inputs.UnpackIntoMap(inputMap, decodedData)
		arr, err := json.Marshal(&inputMap)
		if err != nil {
			return err
		}
		// 反序列化
		err = json.Unmarshal(arr, &functionInput)
		if err != nil {
			return err
		}
		contractAddress = strings.ToLower(functionInput.TokenAddress[0].String())
		in1CoinAddress := false
		if _, ok := coinAddress[contractAddress]; ok {
			in1CoinAddress = true
		}
		if in1CoinAddress == false {
			return nil
		}
		amount = functionInput.Amount[0].Div(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals)))).String()
		coinToken = coinAddress[contractAddress].Name
		if len(functionInput.TokenAddress) > 1 {
			contractAddress1 = strings.ToLower(functionInput.TokenAddress[1].String())
			in2CoinAddress := false
			if _, ok := coinAddress[contractAddress1]; ok {
				in2CoinAddress = true
			}
			if in2CoinAddress == false {
				return nil
			}
			amount1 = functionInput.Amount[1].Div(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress1].Decimals)))).String()
			coinToken1 = coinAddress[contractAddress1].Name
		}
		rechargeType = int8(len(functionInput.TokenAddress))
		remarks = functionInput.Remarks

		customeAmount1, _ := json.Marshal(functionInput.CustomeAmount)
		customeAmount = gconv.String(customeAmount1)

		customeCoin1, _ := json.Marshal(functionInput.CustomeCoin)
		customeCoin = gconv.String(customeCoin1)

		customeUser1, _ := json.Marshal(functionInput.CustomeUser)
		customeUser = gconv.String(customeUser1)

	} else if functionName == "0xd737040c" {
		//卡牌合成方法
		reader, err := os.Open("./abi/new_recharge/new_recharge.abi")
		defer reader.Close()
		if err != nil {
			return err
		}
		tokenAbi, err := abi.JSON(reader)
		if err != nil {
			return err
		}
		decodedSig, _ := hex.DecodeString(transfer.Input[2:10])
		method, _ := tokenAbi.MethodById(decodedSig)
		decodedData, _ := hex.DecodeString(transfer.Input[10:])
		var functionInput rpc.CardSynthesisCardInputs
		inputMap := make(map[string]interface{}, 0)
		_ = method.Inputs.UnpackIntoMap(inputMap, decodedData)
		arr, err := json.Marshal(&inputMap)
		if err != nil {
			return err
		}
		// 反序列化
		err = json.Unmarshal(arr, &functionInput)
		if err != nil {
			return err
		}
		contractAddress = "卡牌合成"
		amount = "0"
		coinToken = "卡牌合成"
		rechargeType = 1
		remarks = functionInput.Remarks
	} else if functionName == "0xfad3cc4b" {
		//dapp-horc项目特殊情况
		reader, err := os.Open("./abi/dappHorc/dappHorc.json")
		defer reader.Close()
		if err != nil {
			return err
		}
		contractAbi, err := abi.JSON(reader)
		if err != nil {
			return err
		}
		decodedSig, _ := hex.DecodeString(transfer.Input[2:10])
		method, _ := contractAbi.MethodById(decodedSig)
		decodedData, _ := hex.DecodeString(transfer.Input[10:])
		var functionInput rpc.DappHorcDepositInputs
		inputMap := make(map[string]interface{}, 0)
		_ = method.Inputs.UnpackIntoMap(inputMap, decodedData)
		arr, err := json.Marshal(&inputMap)
		if err != nil {
			return err
		}
		// 反序列化
		err = json.Unmarshal(arr, &functionInput)
		if err != nil {
			return err
		}
		contractAddress = "dapp-horc质押"
		amount = decimal.NewFromBigInt(functionInput.Amount, 0).DivRound(decimal.NewFromFloat(math.Pow(10, 18)), 18).String()
		coinToken = "dapp-horc质押"
		rechargeType = 1
		remarks = functionInput.Invitor.String()
	} else if functionName == "0x4e71d92d" {
		contractAddress = "dapp-horc质押"
		amount = "无法获取"
		coinToken = "dapp-horc领取收益"
		rechargeType = 1
		remarks = "claim"
	} else if functionName == "0x1afce229" {
		contractAbi, err := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address[]","name":"tokenAddress","type":"address[]"},{"indexed":false,"internalType":"uint256[]","name":"amount","type":"uint256[]"},{"indexed":false,"internalType":"string","name":"remark","type":"string"},{"indexed":false,"internalType":"address[]","name":"customerAddress","type":"address[]"}],"name":"TransactionDetails","type":"event"}]`))
		if err != nil {
			return err
		}
		var event struct {
			TokenAddress    []common.Address
			Amount          []*big.Int
			Remark          string
			CustomerAddress []common.Address
		}
		for _, vlog := range bscStruct.Logs {
			if len(vlog.Topics) > 0 && vlog.Topics[0] == contractAbi.Events["TransactionDetails"].ID.String() {
				decodedSig, err := hex.DecodeString(vlog.Data[2:])
				if err != nil {
					return err
				}
				err = contractAbi.UnpackIntoInterface(&event, "TransactionDetails", decodedSig)
				if err != nil {
					return err
				}

				contractAddress = strings.ToLower(event.TokenAddress[0].String())
				coinToken = coinAddress[contractAddress].Name
				amount = decimal.NewFromBigInt(event.Amount[0], 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
				if len(event.TokenAddress) >= 2 {
					contractAddress1 = strings.ToLower(event.TokenAddress[1].String())
					coinToken1 = coinAddress[contractAddress1].Name
					amount1 = decimal.NewFromBigInt(event.Amount[1], 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress1].Decimals))), 18).String()
				}
				remarks = event.Remark
				rechargeType = 1
			}
		}
	} else {
		return nil
	}
	toAddress := transfer.From
	status = 3

	//检查如果hash不存在，则可新增
	existsRecharge, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if existsRecharge == nil {
		data1 := dao.RechargeAddReq{MainChain: "bsc", BlockHash: transfer.BlockHash, CoinToken: coinToken, CoinToken1: coinToken1, FromAddress: fromAddress, ToAddress: toAddress, Amount: amount, Amount1: amount1, ContractAddress: contractAddress,
			ContractAddress1: contractAddress1, Hash: transfer.Hash, BlockHeight: blockHeight, Status: status, Remarks: remarks, RechargeType: rechargeType, TokenId: tokenId,
			CustomeUser: customeUser, CustomeCoin: customeCoin, CustomeAmount: customeAmount}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		sendNotify(&data1)
	}
	return nil
}

// ERC20 Transfer事件的签名哈希
const ERC20_TRANSFER_TOPIC = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

func addressRechargeHandle(transfer *rpc.BscTransactions) (err error) {
	ctx := gctx.New()
	userAddress, err := sservice.Address.GetBnbAllAddress(ctx)
	if err != nil {
		return err
	}
	coinAddress, err := sservice.Currency.GetBnbCoinAddress(ctx)
	if err != nil {
		return err
	}

	//检查交易状态
	client := rpc.BscClient{}
	data, err := client.Init().GetTransferStatus(transfer.Hash)
	if err != nil {
		g.Log().Printf("查询交易%v失败，%v# \n", transfer.Hash, err)
		return err
	}
	bscStruct := rpc.BscTransferDetail{}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		g.Log().Printf("查询交易%v失败，%v# \n", transfer.Hash, err)
		return err
	}
	if bscStruct.Status == "0x0" {
		g.Log().Printf("交易状态是失败的，不处理，%v# \n", transfer.Hash, err)
		return nil
	}

	// 检查是否是直接的主币转账（BNB）
	isMainCoinTransfer := transfer.Input == "0x"
	if isMainCoinTransfer {
		return handleMainCoinTransfer(transfer, userAddress, coinAddress, &bscStruct)
	}

	// 检查是否是通过Transfer事件的代币转账
	transferEvents := parseTransferEvents(bscStruct.Logs, userAddress, coinAddress)
	if len(transferEvents) > 0 {
		return handleTokenTransferByEvents(transfer, transferEvents, &bscStruct)
	}

	// 兼容原有的直接ERC20 transfer调用逻辑
	functionName := library.SubStr(transfer.Input, 0, 10)
	if functionName == "0xa9059cbb" {
		return handleDirectTokenTransfer(transfer, userAddress, coinAddress, &bscStruct)
	}

	return nil
}

// TransferEvent 代币转账事件信息
type TransferEvent struct {
	ContractAddress string
	FromAddress     string
	ToAddress       string
	Amount          string
	CoinToken       string
	Decimals        int
}

// parseTransferEvents 解析Transfer事件日志，返回匹配的转账事件
func parseTransferEvents(logs []rpc.BscTransferDetailLog, userAddressMap []string, coinAddressMap map[string]*model.Currency) []TransferEvent {
	var transferEvents []TransferEvent
	userAddressSet := make(map[string]bool)
	for _, addr := range userAddressMap {
		userAddressSet[strings.ToLower(addr)] = true
	}

	for _, log := range logs {
		// 检查是否为Transfer事件
		if len(log.Topics) >= 3 && strings.ToLower(log.Topics[0]) == strings.ToLower(ERC20_TRANSFER_TOPIC) {
			// Transfer事件格式: Transfer(address indexed from, address indexed to, uint256 value)
			// topics[0]: 事件签名
			// topics[1]: from地址
			// topics[2]: to地址
			// data: 转账金额

			contractAddress := strings.ToLower(log.Address)

			// 检查是否为我们关注的币种
			coinInfo, isValidCoin := coinAddressMap[contractAddress]
			if !isValidCoin {
				continue
			}

			// 解析from地址 (去掉0x前缀，取后40位，补0x前缀)
			fromAddressTopic := log.Topics[1]
			fromAddress := "0x" + fromAddressTopic[len(fromAddressTopic)-40:]

			// 解析to地址
			toAddressTopic := log.Topics[2]
			toAddress := "0x" + toAddressTopic[len(toAddressTopic)-40:]
			toAddressLower := strings.ToLower(toAddress)

			// 检查是否转给我们关注的用户地址
			if userAddressSet[toAddressLower] {
				// 解析转账金额
				amountHex := log.Data
				if amountHex == "0x" || amountHex == "" {
					continue
				}
				amount := decimal.NewFromBigInt(library.HexToBigInt(amountHex), int32(-coinInfo.Decimals)).Truncate(int32(coinInfo.Decimals)).String()

				transferEvent := TransferEvent{
					ContractAddress: contractAddress,
					FromAddress:     strings.ToLower(fromAddress),
					ToAddress:       toAddressLower,
					Amount:          amount,
					CoinToken:       coinInfo.Name,
					Decimals:        coinInfo.Decimals,
				}
				transferEvents = append(transferEvents, transferEvent)
			}
		}
	}

	return transferEvents
}

// handleMainCoinTransfer 处理主币（BNB）转账
func handleMainCoinTransfer(transfer *rpc.BscTransactions, userAddress []string, coinAddress map[string]*model.Currency, bscStruct *rpc.BscTransferDetail) error {
	ctx := gctx.New()
	toAddress := strings.ToLower(transfer.To)

	// 检查是否转给我们关注的地址
	isValidToAddress := false
	for _, addr := range userAddress {
		if strings.ToLower(addr) == toAddress {
			isValidToAddress = true
			break
		}
	}
	if !isValidToAddress {
		return nil
	}

	// BNB的合约地址标识
	contractAddress := "0x0000000000000000000000000000000000000000"
	coinInfo, exists := coinAddress[contractAddress]
	if !exists {
		return nil
	}

	fromAddress := transfer.From
	coinToken := "BNB"
	amount := decimal.NewFromBigInt(library.HexToBigInt(transfer.Value), int32(-coinInfo.Decimals)).Truncate(int32(coinInfo.Decimals)).String()
	status := int8(1)
	rechargeType := int8(1)
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)

	//检查如果hash不存在，则可新增
	existsRecharge, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if existsRecharge == nil {
		data1 := dao.RechargeAddReq{
			MainChain:       "bsc",
			BlockHash:       transfer.BlockHash,
			CoinToken:       coinToken,
			FromAddress:     fromAddress,
			ToAddress:       toAddress,
			Amount:          amount,
			ContractAddress: contractAddress,
			Hash:            transfer.Hash,
			BlockHeight:     blockHeight,
			RechargeType:    rechargeType,
			Status:          status,
		}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		sendNotify(&data1)
	}
	return nil
}

// handleTokenTransferByEvents 处理通过事件检测到的代币转账
func handleTokenTransferByEvents(transfer *rpc.BscTransactions, transferEvents []TransferEvent, bscStruct *rpc.BscTransferDetail) error {
	ctx := gctx.New()
	fromAddress := transfer.From
	status := int8(1)
	rechargeType := int8(1)
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)

	// 处理每个检测到的转账事件
	for _, event := range transferEvents {
		//检查如果hash不存在，则可新增
		existsRecharge, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
		if err != nil {
			return err
		}
		if existsRecharge == nil {
			data1 := dao.RechargeAddReq{
				MainChain:       "bsc",
				BlockHash:       transfer.BlockHash,
				CoinToken:       event.CoinToken,
				FromAddress:     fromAddress,
				ToAddress:       event.ToAddress,
				Amount:          event.Amount,
				ContractAddress: event.ContractAddress,
				Hash:            transfer.Hash,
				BlockHeight:     blockHeight,
				RechargeType:    rechargeType,
				Status:          status,
			}
			err = sservice.Recharge.Add(ctx, &data1)
			if err != nil {
				g.Log().Printf("插入交易失败 %v \n", err)
				return err
			}
			sendNotify(&data1)
		}
		// 只处理第一个匹配的转账事件，避免重复记录
		break
	}
	return nil
}

// handleDirectTokenTransfer 处理直接的ERC20 transfer调用（兼容原有逻辑）
func handleDirectTokenTransfer(transfer *rpc.BscTransactions, userAddress []string, coinAddress map[string]*model.Currency, bscStruct *rpc.BscTransferDetail) error {
	ctx := gctx.New()
	contractAddress := strings.ToLower(transfer.To)

	// 检查是否为我们关注的币种
	coinInfo, exists := coinAddress[contractAddress]
	if !exists {
		return nil
	}

	// 解析ERC20 transfer函数参数
	toAddress := "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, -128, 64), "0"), 40, "0")

	// 检查是否转给我们关注的地址
	isValidToAddress := false
	for _, addr := range userAddress {
		if strings.ToLower(addr) == strings.ToLower(toAddress) {
			isValidToAddress = true
			break
		}
	}
	if !isValidToAddress {
		return nil
	}

	fromAddress := transfer.From
	coinToken := coinInfo.Name
	amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, -64, -1), "0")
	if amountString == "0x" {
		return nil
	}
	amount := decimal.NewFromBigInt(library.HexToBigInt(amountString), int32(-coinInfo.Decimals)).Truncate(int32(coinInfo.Decimals)).String()
	status := int8(1)
	rechargeType := int8(1)
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)

	//检查如果hash不存在，则可新增
	existsRecharge, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if existsRecharge == nil {
		data1 := dao.RechargeAddReq{
			MainChain:       "bsc",
			BlockHash:       transfer.BlockHash,
			CoinToken:       coinToken,
			FromAddress:     fromAddress,
			ToAddress:       strings.ToLower(toAddress),
			Amount:          amount,
			ContractAddress: contractAddress,
			Hash:            transfer.Hash,
			BlockHeight:     blockHeight,
			RechargeType:    rechargeType,
			Status:          status,
		}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		sendNotify(&data1)
	}
	return nil
}

func sendNotify(recharge *dao.RechargeAddReq) {
	//检测是否配置了充值回调地址
	callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
	if callbackUrl.ConfigValue == "" {
		g.Log().File("callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v \n", recharge.Hash)
		return
	}

	// 处理 remarks 去除空字节
	cleanRemarks := strings.ReplaceAll(recharge.Remarks, "\x00", "")

	data := url.Values{
		"main_chain":        {recharge.MainChain},
		"block_hash":        {recharge.BlockHash},
		"recharge_type":     {gconv.String(recharge.RechargeType)},
		"from_address":      {recharge.FromAddress},
		"to_address":        {recharge.ToAddress},
		"coin_token":        {recharge.CoinToken},
		"coin_token1":       {recharge.CoinToken1},
		"contract_address":  {recharge.ContractAddress},
		"contract_address1": {recharge.ContractAddress1},
		"amount":            {recharge.Amount},
		"amount1":           {gconv.String(recharge.Amount1)},
		"hash":              {recharge.Hash},
		"imputation_hash":   {""},
		"remarks":           {cleanRemarks},
		"status":            {gconv.String(recharge.Status)},
		"token_id":          {recharge.TokenId},
		"customeUser":       {recharge.CustomeUser},
		"customeCoin":       {recharge.CustomeCoin},
		"customeAmount":     {recharge.CustomeAmount},
	}
	resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, data)
	defer resp.Body.Close()
	g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, data.Encode(), resp.StatusCode)
	return
}
