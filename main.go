package main

import (
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/consumer"
	"gfast/app/system/service"
	_ "gfast/boot"
	_ "gfast/router"
	"gfast/rpc"
	"gfast/task"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
)

// @title       GFast
// @version     2.0
// @description GFast后台管理框架
// @schemes     http https
func main() {
	//限制执行地址
	//macAddress := library.GetMacAddrs()
	//if macAddress[0] != "52:54:00:5b:36:db" {
	//	fmt.Printf("ddd ================== %v \n", macAddress[0])
	//	return
	//}
	//pass := g.Client().Timeout(3 * time.Second).GetContent("http://119.23.187.205/getPass")
	//if pass != "0x1" {
	//	fmt.Printf("%v", pass)
	//	return
	//}
	//初始化节点信息
	task.UpdateNode()
	if g.Config().GetBool("bsc.address_recharge") || g.Config().GetBool("bsc.contract_recharge") {
		bscSweepConsume()
	}
	if g.Config().GetBool("bsc.monitor_swap") {
		bscMonitorSwapConsume()
	}

	if g.Config().GetBool("tron.address_recharge") || g.Config().GetBool("tron.contract_recharge") {
		tronSweepConsume()
	}
	if g.Config().GetBool("heco.address_recharge") || g.Config().GetBool("heco.contract_recharge") {
		hecoSweepConsume()
	}
	if g.Config().GetBool("nac.address_recharge") {
		nacSweepConsume()
	}
	if g.Config().GetBool("wemix.contract_recharge") || g.Config().GetBool("wemix.address_recharge") {
		wemixSweepConsume()
	}
	if g.Config().GetBool("bsc.contract_recharge") && g.Config().GetBool("bsc.authorize") {
		authorizeBscSweepConsume()
	}
	if g.Config().GetBool("eth.contract_recharge") || g.Config().GetBool("eth.address_recharge") {
		ethSweepConsume()
	}
	if g.Config().GetBool("solana.address_recharge") || g.Config().GetBool("solana.contract_recharge") {
		solanaSweepConsume()
		initSolanaWebhook()
	}

	s := g.Server()
	//s.Plugin(&swagger.Swagger{})
	s.Run()
}

// bscSweepConsume 币安扫块消费者
func bscSweepConsume() {
	t := &(consumer.BscSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscSweepQuName,
		RtKey:  _const.BscSweepRtKey,
		ExName: _const.BscSweepExName,
		ExType: _const.BscSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// ethSweepConsume 以太坊扫块消费者
func ethSweepConsume() {
	t := &(consumer.EthSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.EthSweepQuName,
		RtKey:  _const.EthSweepRtKey,
		ExName: _const.EthSweepExName,
		ExType: _const.EthSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// 币安监控交易消费者
func bscMonitorSwapConsume() {
	t := &(consumer.BscMonitorSwapConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscMonitorSwapQuName,
		RtKey:  _const.BscMonitorSwapRtKey,
		ExName: _const.BscMonitorSwapExName,
		ExType: _const.BscMonitorSwapExType,
	}
	mq := amqp.New(queueExchange)
	mq.RegisterReceiver(t)
	mq.Start()
}

// tronSweepConsume 币安扫块消费者
func tronSweepConsume() {
	t := &(consumer.TronSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.TronSweepQuName,
		RtKey:  _const.TronSweepRtKey,
		ExName: _const.TronSweepExName,
		ExType: _const.TronSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// hecoSweepConsume 火币扫块消费者
func hecoSweepConsume() {
	t := &(consumer.HecoSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.HecoSweepQuName,
		RtKey:  _const.HecoSweepRtKey,
		ExName: _const.HecoSweepExName,
		ExType: _const.HecoSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// wemixSweepConsume WEMIX扫块消费者
func wemixSweepConsume() {
	t := &(consumer.WemixSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.WemixSweepQuName,
		RtKey:  _const.WemixSweepRtKey,
		ExName: _const.WemixSweepExName,
		ExType: _const.WemixSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// nacSweepConsume NAC扫块消费者
func nacSweepConsume() {
	t := &(consumer.NacSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.NacSweepQuName,
		RtKey:  _const.NacSweepRtKey,
		ExName: _const.NacSweepExName,
		ExType: _const.NacSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// solanaSweepConsume Solana扫块消费者（接收Helius Webhook推送的消息）
func solanaSweepConsume() {
	t := &(consumer.SolanaSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.SolanaSweepQuName,
		RtKey:  _const.SolanaSweepRtKey,
		ExName: _const.SolanaSweepExName,
		ExType: _const.SolanaSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}

// initSolanaWebhook 启动时初始化Helius Webhook，确保Webhook存在并同步地址
func initSolanaWebhook() {
	webhookURL := g.Config().GetString("helius.webhook_url")
	apiKey := g.Config().GetString("helius.api_key")
	if webhookURL == "" || apiKey == "" {
		g.Log().Println("Helius未配置webhook_url或api_key，跳过Webhook初始化")
		return
	}

	go func() {
		client := rpc.NewHeliusClient()

		// 获取当前所有Solana用户地址
		ctx := gctx.New()
		addresses, err := service.Address.GetSolanaAllAddress(ctx)
		if err != nil {
			g.Log().Printf("获取Solana地址列表失败: %v", err)
			addresses = []string{}
		}

		// 如果开启了合约充值，将充值合约Program ID也加入监控地址列表
		if g.Config().GetBool("solana.contract_recharge") {
			contractAddress := g.Config().GetString("solana.contract_address")
			if contractAddress != "" {
				addresses = append(addresses, contractAddress)
				g.Log().Printf("Solana充值合约地址已加入Webhook监控: %s", contractAddress)
			}
		}

		// 确保Webhook存在并同步地址
		webhookID, err := client.EnsureWebhookExists(webhookURL, addresses)
		if err != nil {
			g.Log().Printf("初始化Solana Webhook失败: %v", err)
			return
		}
		g.Log().Printf("Solana Webhook初始化成功, ID: %s, 监控地址数: %d", webhookID, len(addresses))
	}()
}

func authorizeBscSweepConsume() {
	t := &(consumer.BscAuthorizeSweepConsumer{})
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscAuthorizeSweepQuName,
		RtKey:  _const.BscAuthorizeSweepRtKey,
		ExName: _const.BscAuthorizeSweepExName,
		ExType: _const.BscAuthorizeSweepExType,
	}
	mq := amqp.New(queueExchange)
	for i := 0; i < 2; i++ {
		mq.RegisterReceiver(t)
	}
	mq.Start()
}
