package consumer

import (
	"encoding/json"
	"gfast/abi/erc20"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	systemService "gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

type BscAuthorizeSweepConsumer struct{}

// Consumer 实现币安扫块消费者
func (b *BscAuthorizeSweepConsumer) Consumer(dataByte []byte, key uint64) error {
	transfer := rpc.BscTransactions{}
	err := json.Unmarshal(dataByte, &transfer)
	if err != nil {
		return err
	} //json解析到结构体里面
	g.Log().File("authorize-consumer.{Y-m-d}.log").Printf("消费 %v", transfer.Hash)
	//判断一下这是合约充值还是获取地址充值
	err = authorizeContractRechargeHandle(&transfer)
	return err
}

func authorizeContractRechargeHandle(transfer *rpc.BscTransactions) (err error) {
	//检查交易状态
	client := rpc.BscClient{}
	data, err := client.Init().GetTransferStatus(transfer.Hash)
	if err != nil {
		return err
	}
	bscStruct := rpc.BscTransferDetail{}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		return err
	}
	if bscStruct.Status == "0x0" {
		return nil
	}
	ctx := gctx.New()
	if err != nil {
		return err
	}

	functionName := library.SubStr(transfer.Input, 0, 10)
	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)
	authorizeHash := transfer.Hash
	var (
		contractAddress string
		address         string
		coinName        string
		coinDecimals    uint8
		coinAddress     string
		num             string
	)

	if functionName == "0x095ea7b3" {
		//合约地址
		contractAddress = "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(transfer.Input, 10, 64), "0"), 40, "0")
		address = transfer.From
		coinAddress = transfer.To

		//查询代币信息
		cache := service.Cache.New()
		rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
		client, _ := ethclient.Dial(rpcUrl)
		erc20Contract, err := erc20.NewErc20AbiCaller(common.HexToAddress(coinAddress), client)
		if err != nil {
			return err
		}
		transactOpts := &bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}
		coinName, _ = erc20Contract.Name(transactOpts)
		coinDecimals, _ = erc20Contract.Decimals(transactOpts)

		amountString := "0x" + strings.TrimLeft(library.SubStr(transfer.Input, 74, 64), "0")
		if amountString == "0x" {
			return nil
		}
		num = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinDecimals))), int32(coinDecimals)).Truncate(18).String()
	}

	//检查如果hash不存在，则可新增
	exists, err := systemService.AuthorizeAddress.GetInfoByHash(ctx, transfer.Hash)
	if err != nil {
		return err
	}
	if exists == nil {
		data1 := dao.AuthorizeAddressAddReq{MainChain: "bsc", ContractAddress: contractAddress, Address: address, CoinName: coinName,
			CoinDecimals: gconv.Int(coinDecimals), CoinAddress: coinAddress, Num: num, AuthorizeHash: authorizeHash, AuthorizeBlock: blockHeight}
		err = systemService.AuthorizeAddress.Add(ctx, &data1)
		if err != nil {
			return err
		}
	}
	return nil
}
