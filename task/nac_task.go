package task

import (
	"encoding/json"
	"fmt"
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
)

func nacSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("nac.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache    = service.Cache.New()
		nowBlock int64
		newBlock int64
		err      error
	)
	cacheBlock := cache.Get(global.NacSweepBlockNumber)
	if cacheBlock == nil {
		nowBlock, err = rpc.GetLastHeight()
		if err != nil {
			g.Log().Printf("NAC获取新快错误%v", err)
			return
		}
		cache.Set(global.NacSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	newBlock, _ = rpc.GetLastHeight()
	g.Log().Printf("NAC最新块%v", newBlock)
	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetNacAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetNacCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.NacSweepQuName,
			RtKey:  _const.NacSweepRtKey,
			ExName: _const.NacSweepExName,
			ExType: _const.NacSweepExType,
		}
		mq := amqp.New(queueExchange)
		var current int64
		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("NAC扫描块%v", current)
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
			cache.Set(global.NacSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
