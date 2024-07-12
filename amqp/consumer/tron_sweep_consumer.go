package consumer

import (
	"encoding/hex"
	"encoding/json"
	"gfast/app/common/model"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	sservice "gfast/app/system/service"
	"gfast/hdwallet"
	"gfast/library"
	"gfast/rpc"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"math"
	"net/url"
	"strings"
)

type TronSweepConsumer struct{}

// Consumer 实现波场扫块消费者
func (b *TronSweepConsumer) Consumer(dataByte []byte, key uint64) error {
	tronTransaction := &rpc.TronTransaction{}
	err := json.Unmarshal(dataByte, &tronTransaction)
	if err != nil {
		return err
	}

	value := api.TransactionExtention{}
	err = proto.Unmarshal(tronTransaction.Transaction, &value)
	if err != nil {
		return err
	}

	g.Log().File("tron-consumer.{Y-m-d}.log").Printf("消费 %v", hex.EncodeToString(value.Txid))

	if value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TriggerSmartContract {
		err = tronContractRechargeHandle(&value, tronTransaction.BlockHeader, tronTransaction.BlockHash)
	}

	if value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TransferContract {
		err = tronAddressRechargeHandle(&value, tronTransaction.BlockHeader, tronTransaction.BlockHash)
	}
	return err
}

// 这里有两种情况会进来，一种是合约充值，一种是对应代币的交易会进来
func tronContractRechargeHandle(transfer *api.TransactionExtention, blockHeight int64, blockHash string) (err error) {
	//判断如果是手续费地址，或者提币地址，直接跳过
	unObj := &core.TriggerSmartContract{}
	err = proto.Unmarshal(transfer.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj)
	if err != nil {
		return err
	}

	fromAddress := hdwallet.EncodeCheck(unObj.GetOwnerAddress())
	contractAddress := hdwallet.EncodeCheck(unObj.GetContractAddress())
	ctx := gctx.New()
	coinAddress, err := sservice.Currency.GetTronCoinAddress(ctx)
	if err != nil {
		return err
	}
	input := hex.EncodeToString(unObj.GetData())
	functionName := library.SubStr(input, 0, 8)
	transferHash := hex.EncodeToString(transfer.GetTxid())
	var (
		amount           string
		coinToken        string
		remarks          string
		rechargeType     int8
		amount1          string
		coinToken1       string
		contractAddress1 string
		status           int8
		inCoinAddress    bool
		toAddress        string
	)

	//代币转账
	if functionName == "a9059cbb" {
		//单币充值
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(input, 72, 64), "0")
		if amountString == "0x" {
			return nil
		}
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		coinToken = coinAddress[contractAddress].Name
		rechargeType = 1
		toAddress, _ = hdwallet.FromHexAddress("41" + input[32:72])
		status = 1
	} else if functionName == "23b872dd" {
		//授权转账
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(input, 136, -1), "0")
		if amountString == "0x" {
			return nil
		}
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		coinToken = coinAddress[contractAddress].Name
		rechargeType = 1
		toAddress, _ = hdwallet.FromHexAddress("41" + input[96:136])
		status = 1
	} else if functionName == "68ca0399" {
		//单币充值
		contractAddress, _ = hdwallet.FromHexAddress("41" + library.StrPadLeft(strings.TrimLeft(library.SubStr(input, 8, 64), "0"), 40, "0"))
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amountString := "0x" + strings.TrimLeft(library.SubStr(input, 72, 64), "0")
		if amountString == "0x" {
			return nil
		}
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		coinToken = coinAddress[contractAddress].Name
		fsd, err := hex.DecodeString(input[264:])
		if err != nil {
			return err
		}
		remarks = gconv.String(fsd)
		rechargeType = 1
		toAddress = hdwallet.EncodeCheck(unObj.GetOwnerAddress())
		status = 3
	} else if functionName == "34fe59b1" {
		//充值卡牌
		contractAddress, _ = hdwallet.FromHexAddress("41" + library.StrPadLeft(strings.TrimLeft(library.SubStr(input, 8, 64), "0"), 40, "0"))
		if _, ok := coinAddress[contractAddress]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			return nil
		}
		amount = "1"
		coinToken = coinAddress[contractAddress].Name
		fsd, err := hex.DecodeString(input[136:])
		if err != nil {
			return err
		}
		remarks = gconv.String(fsd)
		rechargeType = 1
		toAddress = hdwallet.EncodeCheck(unObj.GetOwnerAddress())
		status = 3
	} else if functionName == "a6dba280" {
		//双币充值
		contractAddress, _ = hdwallet.FromHexAddress("41" + library.StrPadLeft(strings.TrimLeft(library.SubStr(input, 8, 64), "0"), 40, "0"))
		contractAddress1, _ = hdwallet.FromHexAddress("41" + library.StrPadLeft(strings.TrimLeft(library.SubStr(input, 136, 64), "0"), 40, "0"))
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
		amountString := "0x" + strings.TrimLeft(library.SubStr(input, 72, 64), "0")
		amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18).String()
		amount1String := "0x" + strings.TrimLeft(library.SubStr(input, 64*3+8, 64), "0")
		amount1 = decimal.NewFromBigInt(library.HexToBigInt(amount1String), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress1].Decimals))), 18).String()
		rechargeType = 2
		coinToken = coinAddress[contractAddress].Name
		coinToken1 = coinAddress[contractAddress1].Name
		fsd, err := hex.DecodeString(input[392:])
		if err != nil {
			return err
		}
		remarks = gconv.String(fsd)
		toAddress = hdwallet.EncodeCheck(unObj.GetOwnerAddress())
		status = 3
	} else {
		//其他情况直接不考虑
		return nil
	}

	//检查如果hash不存在，则可新增
	exists, err := sservice.Recharge.GetInfoByHash(ctx, transferHash)
	if err != nil {
		return err
	}
	if exists == nil {
		data1 := dao.RechargeAddReq{MainChain: "tron", BlockHash: blockHash, CoinToken: coinToken, CoinToken1: coinToken1, FromAddress: fromAddress, ToAddress: toAddress, Amount: amount, Amount1: amount1, ContractAddress: contractAddress,
			ContractAddress1: contractAddress1, Hash: transferHash, BlockHeight: gconv.String(blockHeight), Status: status, Remarks: remarks, RechargeType: rechargeType}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		tronSendNotify(&data1)
	}
	return nil
}

func tronAddressRechargeHandle(transfer *api.TransactionExtention, blockHeight int64, blockHash string) (err error) {
	ctx := gctx.New()
	coinAddress, err := sservice.Currency.GetTronCoinAddress(ctx)
	if err != nil {
		return err
	}

	contractAddress := "TBRop8PopYu8atWWez3g3ueVtSpseW78b6"
	if _, ok := coinAddress[contractAddress]; ok == false {
		return nil
	}
	unObj1 := &core.TransferContract{}
	err = proto.Unmarshal(transfer.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj1)
	if err != nil {
		return err
	}
	transferHash := hex.EncodeToString(transfer.GetTxid())
	fromAddress := hdwallet.EncodeCheck(unObj1.GetOwnerAddress())

	//检查，如果来源地址等于手续费地址，直接退出
	//手续费地址
	var bnbFeeAddress *model.SysConfig
	bnbFeeAddress, err = service.SysConfig.GetConfigByKey("sys.tronFeeAddress")
	if err != nil || bnbFeeAddress == nil {
		return err
	}
	if err == nil && fromAddress == bnbFeeAddress.ConfigValue {
		return nil
	}
	var (
		coinToken    string
		amount       string
		rechargeType int8
		toAddress    string
		status       int8
		remark       string
	)

	toAddress = hdwallet.EncodeCheck(unObj1.GetToAddress())

	coinToken = "TRX"
	amountString := unObj1.GetAmount()
	amount1 := decimal.NewFromInt(amountString).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals))), 18)
	tempAmount, _ := amount1.Float64()
	if tempAmount < 1 {
		return nil
	}
	amount = amount1.String()
	status = 1
	rechargeType = 1
	remark = gconv.String(transfer.Transaction.RawData.Data)
	//检查如果hash不存在，则可新增
	exists, err := sservice.Recharge.GetInfoByHash(ctx, transferHash)
	if err != nil {
		return err
	}
	if exists == nil {
		data1 := dao.RechargeAddReq{MainChain: "tron", BlockHash: blockHash, CoinToken: coinToken, FromAddress: fromAddress, ToAddress: toAddress,
			Amount: amount, ContractAddress: contractAddress, Hash: transferHash, BlockHeight: gconv.String(blockHeight), RechargeType: rechargeType, Status: status, Remarks: remark}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
		tronSendNotify(&data1)
	}
	return nil
}

func tronSendNotify(recharge *dao.RechargeAddReq) {
	//检测是否配置了充值回调地址
	callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
	if callbackUrl.ConfigValue == "" {
		g.Log().Path("data/log/callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v \n", recharge)
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
		"amount1":           {recharge.Amount1},
		"hash":              {recharge.Hash},
		"imputation_hash":   {""},
		"remarks":           {recharge.Remarks},
		"status":            {gconv.String(recharge.Status)},
		"token_id":          {recharge.TokenId},
	}
	resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, data)
	defer resp.Body.Close()
	g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, data.Encode(), resp.StatusCode)
	return
}
