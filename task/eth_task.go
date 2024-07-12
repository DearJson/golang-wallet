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

// 以太坊扫块任务
func ethSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("eth.contract_recharge") && !g.Config().GetBool("eth.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache     = service.Cache.New()
		ethClient = rpc.EthClient{}
		nowBlock  int64
		newBlock  int64
	)
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.EthSweepBlockNumber)
	if cacheBlock == nil {
		rpcBlock, err := ethClient.Init().GetNowBlock()
		if err != nil {
			g.Log().Printf("以太坊获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.EthSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := ethClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)
	g.Log().Printf("以太坊最新块%v", newBlock)

	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetEthAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetEthCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.EthSweepQuName,
			RtKey:  _const.EthSweepRtKey,
			ExName: _const.EthSweepExName,
			ExType: _const.EthSweepExType,
		}
		mq := amqp.New(queueExchange)
		var current int64

		contractRecharge := g.Cfg().GetBool("eth.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("eth.contract_address")
		addressRecharge := g.Cfg().GetBool("eth.address_recharge")
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("以太坊扫描块%v", current)
			blockData, _ := ethClient.Init().GetBlockByNumber(current)
			ethStruct := rpc.EthBlock{}
			err := mapstructure.Decode(blockData, &ethStruct)
			if err != nil {
				g.Log().Printf("以太坊解析块失败，%v %v", current, err)
				return
			}
			for _, value := range ethStruct.Transactions {
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
					t := &(producer.EthProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					g.Log().File("eth-producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}
			}
			cache.Set(global.EthSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
