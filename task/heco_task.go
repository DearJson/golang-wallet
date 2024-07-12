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
func hecoSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("heco.contract_recharge") && !g.Config().GetBool("heco.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache     = service.Cache.New()
		bscClient = rpc.HecoClient{}
		nowBlock  int64
		newBlock  int64
	)
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.HecoSweepBlockNumber)
	if cacheBlock == nil {
		rpcBlock, err := bscClient.Init().GetNowBlock()
		if err != nil {
			g.Log().Printf("火币获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.HecoSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := bscClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)
	g.Log().Printf("火币最新块%v", newBlock)

	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetHecoAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetHecoCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.HecoSweepQuName,
			RtKey:  _const.HecoSweepRtKey,
			ExName: _const.HecoSweepExName,
			ExType: _const.HecoSweepExType,
		}
		mq := amqp.New(queueExchange)
		var current int64

		contractRecharge := g.Cfg().GetBool("heco.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("heco.contract_address")
		addressRecharge := g.Cfg().GetBool("heco.address_recharge")
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("火币扫描块%v", current)
			blockData, _ := bscClient.Init().GetBlockByNumber(current)
			bscStruct := rpc.HecoBlock{}
			err := mapstructure.Decode(blockData, &bscStruct)
			if err != nil {
				g.Log().Printf("火币解析块失败，%v %v", current, err)
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
					t := &(producer.HecoProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					g.Log().File("heco-producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}
			}
			cache.Set(global.HecoSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
