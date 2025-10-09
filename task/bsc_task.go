package task

import (
	"encoding/json"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	"gfast/app/common/service"
	sservice "gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
)

// BSC扫块任务
func bscSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("bsc.contract_recharge") && !g.Config().GetBool("bsc.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache     = service.Cache.New()
		bscClient = rpc.BscClient{}
		nowBlock  int64
		newBlock  int64
	)
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.BscSweepBlockNumber)
	if cacheBlock == nil {
		rpcBlock, err := bscClient.Init().GetNowBlock()
		if err != nil {
			g.Log().Printf("BSC获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.BscSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := bscClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)
	g.Log().Printf("BSC最新块%v", newBlock)

	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetBnbAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetBnbCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.BscSweepQuName,
			RtKey:  _const.BscSweepRtKey,
			ExName: _const.BscSweepExName,
			ExType: _const.BscSweepExType,
		}
		mq := amqp.New(queueExchange)
		var current int64

		contractRecharge := g.Cfg().GetBool("bsc.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("bsc.contract_address")
		addressRecharge := g.Cfg().GetBool("bsc.address_recharge")
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("BSC扫描块%v", current)
			blockData, _ := bscClient.Init().GetBlockByNumber(current)
			BscStruct := rpc.BscBlock{}
			err := mapstructure.Decode(blockData, &BscStruct)
			if err != nil {
				g.Log().Printf("BSC解析块失败，%v %v", current, err)
				return
			}
			for _, value := range BscStruct.Transactions {
				bs := false
				if contractRecharge && value.To == contractRechargeAddress {
					bs = true
				}
				inCoinAddress := false
				if _, ok := coinAddress[value.To]; ok {
					inCoinAddress = true
				}
				if addressRecharge && (library.ElementIsInSlice(value.To, userAddress) || inCoinAddress) {
					bs = true
				}

				if bs {
					jsonByte, _ := json.Marshal(value)
					t := &(producer.BscProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					g.Log().File("Bsc-producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}
			}
			cache.Set(global.BscSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
