package consumer

import (
	"encoding/hex"
	"encoding/json"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	sservice "gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"math"
	"net/url"
	"os"
	"strconv"
	"strings"
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
	if g.Config().GetBool("bsc.contract_recharge") && transfer.To == g.Config().GetString("bsc.contract_address") {
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

	if functionName == "0x68ca0399" {
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
	} else {
		return nil
	}
	toAddress := transfer.From
	status = 3

	//检查如果hash不存在，则可新增
	exists, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if exists == nil {
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
	isContract := transfer.Input != "0x"

	//检查如果是转账代币，如果不是
	functionName := library.SubStr(transfer.Input, 0, 10)
	if isContract && functionName != "0xa9059cbb" {
		return nil
	}
	contractAddress := ""
	if isContract {
		contractAddress = transfer.To
	} else {
		contractAddress = "0x1000000000000000000000000000000000000000"
	}
	if _, ok := coinAddress[contractAddress]; ok == false {
		return nil
	}

	fromAddress := transfer.From
	var (
		coinToken    string
		amount       string
		rechargeType int8
		toAddress    string
		status       int8
	)

	//代币充值
	if isContract {
		toAddress = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, -128, 64), "0"), 40, "0")
		if library.ElementIsInSlice(toAddress, userAddress) == false {
			return nil
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
		coinToken = coinAddress[contractAddress].Name
		amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, -64, -1), "0")
		if amountString == "0x" {
			return nil
		}
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).Div(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals)))).RoundFloor(int32(coinAddress[contractAddress].Decimals)).String()
		//dapp-sus
		//amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, 18)), 18).Mul(decimal.NewFromFloat32(0.97)).Truncate(18).String()
	} else {
		toAddress = transfer.To
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
		coinToken = "BNB"
		amountString := transfer.Value

		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).Truncate(18).String()
	}
	status = 1
	rechargeType = 1
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)
	//检查如果hash不存在，则可新增
	exists, err := sservice.Recharge.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if exists == nil {
		data1 := dao.RechargeAddReq{MainChain: "bsc", BlockHash: transfer.BlockHash, CoinToken: coinToken, FromAddress: fromAddress, ToAddress: toAddress,
			Amount: amount, ContractAddress: contractAddress, Hash: transfer.Hash, BlockHeight: blockHeight, RechargeType: rechargeType, Status: status}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		//if status == 3 {
		sendNotify(&data1)
		//}
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
		"remarks":           {recharge.Remarks},
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
