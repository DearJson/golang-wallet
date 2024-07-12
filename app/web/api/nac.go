package api

import (
	"encoding/json"
	"fmt"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"math/rand"
	"time"
)

type nac struct {
	WebBase
}

var Nac = new(nac)

// GenerateAddress 生成币安地址
func (b *nac) GenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//判断一下该用户标识的bsc链是否存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "nac")
	if exists != nil {
		result := make(map[string]string)
		result["user_id"] = exists.UserId
		result["address"] = exists.Address
		b.SusJsonExit(r, result)
	}
	rpcNewAccount, err := rpc.NewAccount(req.UserId)
	if err != nil {
		b.FailJsonExit(r, "生成失败")
	}
	req.Address = rpcNewAccount.Address
	req.MainChain = "nac"
	req.PrivateKey, _ = library.EncryptByAes(gconv.Bytes(req.PrivateKey))
	err = service.Address.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	result := make(map[string]string)
	result["user_id"] = req.UserId
	result["address"] = req.Address
	b.SusJsonExit(r, result)
}

// Withdraw 提现代币
func (b *nac) Withdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "nac")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	req.MainChain = "nac"
	req.CoinToken = coinInfo.Name
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	b.SusJsonExit(r)
}

// GetBalance 获取余额
func (b *nac) GetBalance(r *ghttp.Request) {
	address := r.GetPost("address")
	tokenId := r.GetPost("tokenId")
	amount, err := rpc.GetBalance(gconv.String(address), gconv.String(tokenId))
	if err != nil {
		b.FailJsonExit(r, "查询余额失败"+err.Error())
	}
	result := make(map[string]string)
	result["balance"] = gconv.String(amount)
	b.SusJsonExit(r, result)
}

func (b *nac) ResetBlock(r *ghttp.Request) {
	ctx := gctx.New()
	userAddress, _ := service.Address.GetNacAllAddress(ctx)
	coinAddress, _ := service.Currency.GetNacCoinAddress(ctx)
	queueExchange := &amqp.QueueExchange{
		QuName: _const.NacSweepQuName,
		RtKey:  _const.NacSweepRtKey,
		ExName: _const.NacSweepExName,
		ExType: _const.NacSweepExType,
	}
	mq := amqp.New(queueExchange)
	current := gconv.Int64(r.GetPost("block_number"))
	nacStruct, err := rpc.GetBlockDetail(current)
	if err != nil {
		g.Log().Printf("NAC解析块失败，%v %v", current, err)
		return
	}
	for _, value := range nacStruct.TxList {
		//查询交易
		txDetail, err := rpc.GetTxDetail(value)
		if err != nil {
			fmt.Printf(" 查询交易失败 == %v \n", err.Error())
		}
		if txDetail.StatusText != "COMPLETED" {
			continue
		}
		inCoinAddress := false
		if _, ok := coinAddress[gconv.String(txDetail.RawTx.Token)]; ok {
			inCoinAddress = true
		}
		if inCoinAddress == false {
			continue
		}
		toAddress := txDetail.ToText
		if library.ElementIsInSlice(toAddress, userAddress) {
			jsonByte, _ := json.Marshal(txDetail)
			t := &(producer.NacProducer{Msg: string(jsonByte)})
			mq.RegisterProducer(t)
			g.Log().File("nac-producer.{Y-m-d}.log").Printf("生产 %v", txDetail.HashText)
		}
	}
	mq.Start()
	b.SusJsonExit(r)
}
