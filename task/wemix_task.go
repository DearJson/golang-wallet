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

// 币安扫块任务
func wemixSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("wemix.contract_recharge") && !g.Config().GetBool("wemix.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache     = service.Cache.New()
		bscClient = rpc.WemixClient{}
		nowBlock  int64
		newBlock  int64
	)
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.WemixSweepBlockNumber)
	if cacheBlock == nil {
		rpcBlock, err := bscClient.Init().GetNowBlock()
		if err != nil {
			g.Log().Printf("wemix获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.WemixSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := bscClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)
	g.Log().Printf("wemix最新块%v", newBlock)

	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetWemixAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetWemixCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.WemixSweepQuName,
			RtKey:  _const.WemixSweepRtKey,
			ExName: _const.WemixSweepExName,
			ExType: _const.WemixSweepExType,
		}
		mq := amqp.New(queueExchange)
		var current int64

		contractRecharge := g.Cfg().GetBool("wemix.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("wemix.contract_address")
		addressRecharge := g.Cfg().GetBool("wemix.address_recharge")
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("wemix扫描块%v", current)
			blockData, _ := bscClient.Init().GetBlockByNumber(current)
			bscStruct := rpc.WemixBlock{}
			err := mapstructure.Decode(blockData, &bscStruct)
			if err != nil {
				g.Log().Printf("wemix解析块失败，%v %v", current, err)
				return
			}
			for _, value := range bscStruct.Transactions {
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
					t := &(producer.WemixProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					g.Log().File("wemix-producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}
			}
			cache.Set(global.WemixSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
