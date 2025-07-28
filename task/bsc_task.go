package task

import (
	"encoding/json"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	"gfast/app/common/service"
	sservice "gfast/app/system/service"
	"gfast/rpc"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
)

// ERC20 Transfer事件的签名哈希
const ERC20_TRANSFER_TOPIC = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

// LogEvent 日志事件结构
type LogEvent struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

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

	ctx := gctx.New()
	userAddress, _ := sservice.Address.GetBnbAllAddress(ctx)
	coinAddress, _ := sservice.Currency.GetBnbCoinAddress(ctx)

	// 优化：将用户地址slice转换为map，提高查找效率从O(n)到O(1)
	userAddressMap := make(map[string]bool, len(userAddress))
	for _, addr := range userAddress {
		if addr != "" {
			userAddressMap[strings.ToLower(addr)] = true
		}
	}

	// 优化：将币种地址map的key转换为小写，统一比较
	coinAddressMap := make(map[string]bool, len(coinAddress))
	coinContractAddresses := make([]string, 0, len(coinAddress))
	for addr := range coinAddress {
		if addr != "" {
			coinAddressMap[strings.ToLower(addr)] = true
			coinContractAddresses = append(coinContractAddresses, addr) // 用于日志查询
		}
	}

	can := newBlock - nowBlock
	if can > 0 {
		if can > 50 {
			can = 50
		}

		// 优化：准备批量获取区块
		blockNumbers := make([]int64, can)
		for i := int64(0); i < can; i++ {
			blockNumbers[i] = nowBlock + i
		}

		fromBlock := blockNumbers[0]
		toBlock := blockNumbers[len(blockNumbers)-1]

		g.Log().Printf("批量获取区块范围: %d - %d (共%d个块)", fromBlock, toBlock, len(blockNumbers))

		// 优化：批量获取所有区块，显著减少RPC请求次数
		blocksData, err := bscClient.Init().GetBlocksByNumbers(blockNumbers)
		if err != nil {
			g.Log().Printf("批量获取区块失败: %v", err)
			return
		}

		queueExchange := &amqp.QueueExchange{
			QuName: _const.BscSweepQuName,
			RtKey:  _const.BscSweepRtKey,
			ExName: _const.BscSweepExName,
			ExType: _const.BscSweepExType,
		}
		mq := amqp.New(queueExchange)

		contractRecharge := g.Cfg().GetBool("bsc.contract_recharge")
		contractRechargeAddress := strings.ToLower(g.Cfg().GetString("bsc.contract_address"))
		addressRecharge := g.Cfg().GetBool("bsc.address_recharge")

		var totalTxProcessed, validTxFound int64 // 统计信息

		// 优化：如果开启了地址充值，批量查询Transfer日志
		validTxHashMap := make(map[string]string) // txHash -> matchReason
		if addressRecharge {
			g.Log().Printf("批量查询Transfer日志: 区块%d-%d", fromBlock, toBlock)
			transferLogs, err := bscClient.Init().GetTransferLogsBatch(fromBlock, toBlock, coinContractAddresses)
			if err != nil {
				g.Log().Printf("批量查询Transfer日志失败: %v", err)
			} else {
				// 解析Transfer日志，找出匹配的交易
				validTxHashMap = parseTransferLogs(transferLogs, userAddressMap, coinAddressMap)
				g.Log().Printf("Transfer日志解析完成，找到%d个匹配的交易", len(validTxHashMap))
			}
		}

		// 处理批量获取的区块数据
		for i, blockData := range blocksData {
			current := blockNumbers[i]
			g.Log().Printf("币安处理块%v", current)

			// 跳过获取失败的区块
			if blockData == nil {
				g.Log().Printf("跳过获取失败的区块%v", current)
				continue
			}

			bscStruct := rpc.BscBlock{}
			err := mapstructure.Decode(blockData, &bscStruct)
			if err != nil {
				g.Log().Printf("币安解析块失败，%v %v", current, err)
				continue // 继续处理下一个块
			}

			blockTxCount := len(bscStruct.Transactions)
			totalTxProcessed += int64(blockTxCount)

			for _, value := range bscStruct.Transactions {
				needProcess := false
				matchReason := ""

				// 处理合约充值：直接检查交易的to地址
				if contractRecharge && value.To != "" {
					valueTo := strings.ToLower(value.To)
					if valueTo == contractRechargeAddress {
						needProcess = true
						matchReason = "contract_recharge"
					}
				}

				// 处理地址充值：检查是否在预先筛选的有效交易中
				if !needProcess {
					if reason, exists := validTxHashMap[value.Hash]; exists {
						needProcess = true
						matchReason = reason
					}
				}

				// 生产有效交易到队列
				if needProcess && isValidTransaction(value) {
					jsonByte, _ := json.Marshal(value)
					t := &(producer.BscProducer{Msg: string(jsonByte)})
					mq.RegisterProducer(t)
					validTxFound++
					g.Log().File("producer.{Y-m-d}.log").Printf("生产 %v (原因:%s)", value.Hash, matchReason)
				}
			}

			// 更新缓存的区块号
			cache.Set(global.BscSweepBlockNumber, current, 0)
		}

		g.Log().Printf("币安扫块完成: 处理%d个块, 总计%d笔交易, 找到%d笔有效交易", can, totalTxProcessed, validTxFound)

		if validTxFound > 0 {
			mq.Start()
		}
	}
}

// parseTransferLogs 解析Transfer日志，返回匹配的交易哈希映射
func parseTransferLogs(transferLogs []interface{}, userAddressMap map[string]bool, coinAddressMap map[string]bool) map[string]string {
	validTxHashMap := make(map[string]string)

	for _, logData := range transferLogs {
		var logEvent LogEvent
		err := mapstructure.Decode(logData, &logEvent)
		if err != nil {
			continue
		}

		// 检查是否为Transfer事件
		if len(logEvent.Topics) >= 3 && strings.ToLower(logEvent.Topics[0]) == ERC20_TRANSFER_TOPIC {
			// Transfer事件格式: Transfer(address indexed from, address indexed to, uint256 value)
			// topics[0]: 事件签名
			// topics[1]: from地址
			// topics[2]: to地址
			// data: 转账金额

			// 解析to地址 (去掉0x前缀，取后40位，补0x前缀)
			toAddressTopic := logEvent.Topics[2]
			if len(toAddressTopic) >= 42 {
				// 从64位hex字符串中提取地址（后40位）
				addressHex := toAddressTopic[len(toAddressTopic)-40:]
				toAddress := "0x" + addressHex
				toAddressLower := strings.ToLower(toAddress)

				// 检查是否为我们关注的用户地址
				if userAddressMap[toAddressLower] {
					validTxHashMap[logEvent.TransactionHash] = "user_transfer_log"
					continue
				}

				// 检查是否为我们关注的币种地址
				if coinAddressMap[toAddressLower] {
					validTxHashMap[logEvent.TransactionHash] = "coin_transfer_log"
					continue
				}
			}
		}
	}

	return validTxHashMap
}

// 优化：增加交易有效性检查
func isValidTransaction(tx rpc.BscTransactions) bool {
	// 检查交易是否有效
	if tx.Hash == "" {
		return false
	}

	// 检查gas价格是否合理（避免异常交易）
	gasPrice := gconv.Int64(tx.GasPrice)
	if gasPrice <= 0 {
		return false
	}

	// 检查交易金额（可以根据需要调整最小金额阈值）
	value := gconv.Int64(tx.Value)
	if value < 0 {
		return false
	}

	return true
}
