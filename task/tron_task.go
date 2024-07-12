/*
* @desc:测试定时任务
* @company:NULL
* @Author: Json   wbjson@gmail.com
* @Date:   2021/7/16 15:52
 */

package task

import (
	"encoding/hex"
	"encoding/json"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	"gfast/app/common/service"
	sservice "gfast/app/system/service"
	"gfast/hdwallet"
	"gfast/library"
	"gfast/rpc"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/protobuf/proto"
)

//波场扫块任务
func tronSweepTask() {
	//判断一下，如果未开启，不继续
	if !g.Config().GetBool("tron.contract_recharge") && !g.Config().GetBool("tron.address_recharge") {
		g.Log().Println("未开启服务,无需执行!!!")
		return
	}
	var (
		cache    = service.Cache.New()
		nowBlock int64
		newBlock int64
	)
	tronClient, err := rpc.NewClient(gconv.String(cache.Get("tron_rpc_url")))
	if err != nil {
		g.Log().Printf("波场节点链接失败！！！%v", err)
		return
	}
	//查询一下是否有缓存的
	cacheBlock := cache.Get(global.TronSweepBlockNumber)
	if cacheBlock == nil {
		rpcData, err := tronClient.GetNowBlock()
		if err != nil {
			g.Log().Printf("波场获取新快错误%v\n", err)
			return
		}
		rpcBlock := rpcData.BlockHeader.GetRawData().Number
		nowBlock = rpcBlock
		cache.Set(global.TronSweepBlockNumber, rpcBlock, 0)
	} else {
		nowBlock = gconv.Int64(cacheBlock)
	}
	rpcData, err := tronClient.GetNowBlock()
	newBlock = rpcData.BlockHeader.GetRawData().Number
	g.Log().Printf("波场最新块%v %v", newBlock, nowBlock)

	can := newBlock - nowBlock
	if can > 0 {
		if can >= 50 {
			can = 50
		}
		var i int64
		ctx := gctx.New()
		userAddress, _ := sservice.Address.GetTronAllAddress(ctx)
		coinAddress, _ := sservice.Currency.GetTronCoinAddress(ctx)
		queueExchange := &amqp.QueueExchange{
			QuName: _const.TronSweepQuName,
			RtKey:  _const.TronSweepQuName,
			ExName: _const.TronSweepQuName,
			ExType: _const.TronSweepQuName,
		}
		mq := amqp.New(queueExchange)
		var (
			current       int64
			bs            bool
			err           error
			data          string
			blockData     *api.BlockExtention
			inCoinAddress bool
			toAddress     string
			method        string
		)
		contractRecharge := g.Cfg().GetBool("tron.contract_recharge")
		contractRechargeAddress := g.Cfg().GetString("tron.contract_address")
		addressRecharge := g.Cfg().GetBool("tron.address_recharge")

		for i = 0; i < can; i++ {
			current = nowBlock + i
			g.Log().Printf("波场扫描块%v\n", current)
			blockData, err = tronClient.GetBlockByNum(current)
			if err != nil {
				g.Log().Printf("波场解析块失败，%v %v", current, err)
				return
			}
			for _, value := range blockData.GetTransactions() {

				if value.Result == nil || !value.Result.Result {
					continue
				}
				rets := value.Transaction.Ret
				if len(rets) < 1 || rets[0].ContractRet != core.Transaction_Result_SUCCESS {
					continue
				}

				bs = false
				inCoinAddress = false
				if value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TriggerSmartContract {
					unObj := &core.TriggerSmartContract{}
					err = proto.Unmarshal(value.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj)
					if err != nil {
						continue
					}
					contractAddress := hdwallet.EncodeCheck(unObj.GetContractAddress())

					if contractRecharge && contractAddress == contractRechargeAddress {
						bs = true
					}

					if _, ok := coinAddress[contractAddress]; ok {
						inCoinAddress = true
					}
					if addressRecharge && inCoinAddress {
						data = hex.EncodeToString(unObj.GetData())
						method = data[0:8]
						if (method == "a9059cbb" || method == "23b872dd") && len(data) >= 136 {
							if method == "a9059cbb" {
								toAddress, _ = hdwallet.FromHexAddress("41" + data[32:72])
							} else {
								toAddress, _ = hdwallet.FromHexAddress("41" + data[96:136])
							}
							if library.ElementIsInSlice(toAddress, userAddress) {
								bs = true
							}
						}
					}
				}

				//这里只有TRX充值才能进来
				if addressRecharge && value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TransferContract {
					unObj1 := &core.TransferContract{}
					err = proto.Unmarshal(value.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj1)
					if err != nil {
						continue
					}

					toAddress := hdwallet.EncodeCheck(unObj1.ToAddress)
					if library.ElementIsInSlice(toAddress, userAddress) {
						bs = true
					}
				}

				if bs {
					aa, _ := proto.Marshal(value)
					msg := rpc.TronTransaction{
						Transaction: aa,
						BlockHeader: blockData.BlockHeader.RawData.Number,
						BlockHash:   hex.EncodeToString(blockData.Blockid),
					}
					dd, _ := json.Marshal(msg)
					t := &(producer.TronProducer{
						Msg: string(dd),
					})
					mq.RegisterProducer(t)
					g.Log().File("tron-producer.{Y-m-d}.log").Printf("生产 %v", hex.EncodeToString(value.Txid))
				}
			}
			cache.Set(global.TronSweepBlockNumber, current, 0)
		}
		mq.Start()
	}
}
