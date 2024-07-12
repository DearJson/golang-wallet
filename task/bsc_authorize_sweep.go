package task

import (
	"encoding/json"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	"gfast/app/common/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"strings"
)

func bscAuthorizeSweepTask() {
	var (
		cache     = service.Cache.New()
		bscClient = rpc.BscClient{}
		nowBlock  int64
		newBlock  int64
	)
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.BscAuthorizeSweepBlockNumber)
	if cacheBlock == nil {
		rpcBlock, err := bscClient.Init().GetNowBlock()
		if err != nil {
			g.Log().Printf("币安获取新快错误%v", err)
			return
		}
		nowBlock = gconv.Int64(rpcBlock)
		cache.Set(global.BscAuthorizeSweepBlockNumber, nowBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcBLock, _ := bscClient.Init().GetNowBlock()
	newBlock = gconv.Int64(rpcBLock)

	//获取bsc链所有合约地址

	can := newBlock - nowBlock
	if can > 0 {
		if can > 1000 {
			can = 1000
		}
		totalEnd := nowBlock + can
		for {
			if nowBlock >= totalEnd {
				break
			}
			go handle(nowBlock, nowBlock+100)
			nowBlock = nowBlock + 100
		}
		cache.Set(global.BscAuthorizeSweepBlockNumber, totalEnd, 0)
	}

}

func handle(start int64, end int64) {
	authorizeContractAddress := g.Config().GetString("bsc.contract_address")
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscAuthorizeSweepQuName,
		RtKey:  _const.BscAuthorizeSweepRtKey,
		ExName: _const.BscAuthorizeSweepExName,
		ExType: _const.BscAuthorizeSweepExType,
	}
	mq := amqp.New(queueExchange)
	var current int64
	var bs bool
	var i int64
	var bscClient = rpc.BscClient{}
	last := end - start
	for i = 0; i < last; i++ {
		current = start + i
		blockData, _ := bscClient.Init().GetBlockByNumber(current)
		bscStruct := rpc.BscBlock{}
		err := mapstructure.Decode(blockData, &bscStruct)
		if err != nil {
			return
		}
		for _, value := range bscStruct.Transactions {
			bs = false
			//如果方法是授权方法，并且input包含合约地址，则达成条件
			if len(value.Input) >= 10 && value.Input[0:10] == "0x095ea7b3" {
				contractAddress := "0x" + library.StrPadLeft(strings.TrimLeft(library.SubStr(value.Input, 10, 64), "0"), 40, "0")
				if contractAddress == authorizeContractAddress {
					bs = true
				}
			}
			if bs {
				jsonByte, _ := json.Marshal(value)
				t := &(producer.BscAuthorizeProducer{Msg: string(jsonByte)})
				mq.RegisterProducer(t)
				g.Log().File("authorize-producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
			}
		}
	}
	mq.Start()
}
