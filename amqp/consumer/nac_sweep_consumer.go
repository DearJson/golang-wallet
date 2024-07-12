package consumer

import (
	"encoding/json"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	sservice "gfast/app/system/service"
	"gfast/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
)

type NacSweepConsumer struct{}

// Consumer 实现NAC扫块消费者
func (b *NacSweepConsumer) Consumer(dataByte []byte, key uint64) error {
	transfer := rpc.NacTxDetail{}
	err := json.Unmarshal(dataByte, &transfer)
	if err != nil {
		return err
	} //json解析到结构体里面
	g.Log().File("nac-consumer.{Y-m-d}.log").Printf("消费 %v", transfer.HashText)
	fromAddress := transfer.FromText
	//手续费地址
	nacFeeAddress, err := service.SysConfig.GetConfigByKey("sys.nacFeeAddress")
	if err != nil {
		return err
	}

	//提币地址
	nacWithdrawAddress, err := service.SysConfig.GetConfigByKey("sys.nacWithdrawAddress")
	if err != nil {
		return err
	}
	if fromAddress == nacWithdrawAddress.ConfigValue || fromAddress == nacFeeAddress.ConfigValue {
		return nil
	}
	err = nacAddressRechargeHandle(&transfer)
	return nil
}

func nacAddressRechargeHandle(transfer *rpc.NacTxDetail) (err error) {
	ctx := gctx.New()
	coinAddress, err := sservice.Currency.GetNacCoinAddress(ctx)
	if err != nil {
		return err
	}
	fromAddress := transfer.FromText
	coinToken := coinAddress[gconv.String(transfer.RawTx.Token)].Name
	amount := transfer.AmountText
	toAddress := transfer.ToText
	blockHeight := gconv.String(transfer.RawTx.BlockHeight)

	exists, err := sservice.Recharge.GetInfoByHash(ctx, transfer.RawTx.Hash)
	if err != nil {
		return err
	}
	if exists == nil {
		data1 := dao.RechargeAddReq{MainChain: "nac", BlockHash: transfer.RawTx.Hash, CoinToken: coinToken, FromAddress: fromAddress, ToAddress: toAddress,
			Amount: amount, ContractAddress: gconv.String(transfer.RawTx.Token), Hash: transfer.RawTx.Hash, BlockHeight: blockHeight, RechargeType: 1, Status: 1}
		err = sservice.Recharge.Add(ctx, &data1)
		if err != nil {
			g.Log().Printf("插入交易失败 %v \n", err)
			return err
		}
	}
	return nil
}
