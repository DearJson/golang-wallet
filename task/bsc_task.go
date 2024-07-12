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
			g.Log().Printf("币安获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.BscSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := bscClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)
	g.Log().Printf("币安最新块%v", newBlock)

	//解析监控交易的地址
	//monitorSwap := g.Cfg().GetBool("bsc.monitor_swap")
	//monitorSwapRouteAddress := g.Cfg().GetString("bsc.monitor_swap_route_address")
	//monitorSwapCoinAddress, _ := service.MonitorCurrency.GetAddress("bsc")

	//liquiditySwap := g.Cfg().GetBool("bsc.liquidity_swap")
	//liquiditySwapRouteAddress := g.Cfg().GetString("bsc.liquidity_swap_route_address")
	//liquiditySwapCoinAddress, _ := service.LiquidityCurrency.GetAddress("bsc")

	ctx := gctx.New()
	userAddress, _ := sservice.Address.GetBnbAllAddress(ctx)
	coinAddress, _ := sservice.Currency.GetBnbCoinAddress(ctx)
	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		queueExchange := &amqp.QueueExchange{
			QuName: _const.BscSweepQuName,
			RtKey:  _const.BscSweepRtKey,
			ExName: _const.BscSweepExName,
			ExType: _const.BscSweepExType,
		}
		mq := amqp.New(queueExchange)

		//monitorSwapExchange := &amqp.QueueExchange{
		//	QuName: _const.BscMonitorSwapQuName,
		//	RtKey:  _const.BscMonitorSwapRtKey,
		//	ExName: _const.BscMonitorSwapExName,
		//	ExType: _const.BscMonitorSwapExType,
		//}
		//monitorMq := amqp.New(monitorSwapExchange)

		//liquiditySwapExchange := &amqp.QueueExchange{
		//	QuName: _const.BscLiquiditySwapQuName,
		//	RtKey:  _const.BscLiquiditySwapRtKey,
		//	ExName: _const.BscLiquiditySwapExName,
		//	ExType: _const.BscLiquiditySwapExType,
		//}
		//liquidityMq := amqp.New(liquiditySwapExchange)

		contractRecharge := g.Cfg().GetBool("bsc.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("bsc.contract_address")
		addressRecharge := g.Cfg().GetBool("bsc.address_recharge")
		var current int64
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("币安扫描块%v", current)
			blockData, _ := bscClient.Init().GetBlockByNumber(current)
			bscStruct := rpc.BscBlock{}
			err := mapstructure.Decode(blockData, &bscStruct)
			if err != nil {
				g.Log().Printf("币安解析块失败，%v %v", current, err)
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
					t := &(producer.BscProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					g.Log().File("producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}

				//监控swap交易
				//if monitorSwap && monitorSwapRouteAddress == value.To {
				//	for _, monitorCoinAddress := range monitorSwapCoinAddress {
				//		if strings.Contains(value.Input, monitorCoinAddress[2:]) {
				//			jsonByte, _ := json.Marshal(value)
				//			t := &(producer.BscMonitorSwapProducer{Msg: string(jsonByte)})
				//			monitorMq.RegisterProducer(t)
				//			g.Log().File("monitor_producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				//		}
				//	}
				//}

				//监控流动性交易
				//if liquiditySwap {
				//	if liquiditySwapRouteAddress == value.To {
				//		for _, liquidityCoinAddress := range liquiditySwapCoinAddress {
				//			if strings.Contains(value.Input, liquidityCoinAddress[2:]) {
				//				jsonByte, _ := json.Marshal(value)
				//				t := &(producer.BscLiquiditySwapProducer{Msg: string(jsonByte)})
				//				liquidityMq.RegisterProducer(t)
				//				g.Log().File("liquidity_producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				//			}
				//		}
				//	}
				//	//也有可能转移流动性，也需要监听一下
				//	//if library.ElementIsInSlice(value.To, liquiditySwapCoinAddress) {
				//	//	jsonByte, _ := json.Marshal(value)
				//	//	t := &(producer.BscLiquiditySwapProducer{Msg: string(jsonByte)})
				//	//	liquidityMq.RegisterProducer(t)
				//	//	g.Log().File("liquidity_producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				//	//}
				//}
			}
			cache.Set(global.BscSweepBlockNumber, current, 0)
		}
		mq.Start()
		//monitorMq.Start()
	}
}
