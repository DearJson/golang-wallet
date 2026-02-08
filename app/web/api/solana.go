package api

import (
	"encoding/json"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/hdwallet"
	"gfast/library"
	"gfast/rpc"
	"math/rand"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/shopspring/decimal"
)

type solana struct {
	WebBase
}

var Solana = new(solana)

// GenerateAddress 生成Solana地址
func (s *solana) GenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	if err := r.Parse(&req); err != nil {
		s.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	// 判断该用户标识的solana链是否已存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "solana")
	if exists != nil {
		result := make(map[string]string)
		result["user_id"] = exists.UserId
		result["address"] = exists.Address
		s.SusJsonExit(r, result)
	}

	// 生成Ed25519密钥对
	privateKeyHex, address, err := hdwallet.GenerateSolanaKeyPair()
	if err != nil {
		s.FailJsonExit(r, "生成Solana地址失败")
	}

	// 加密私钥存储
	req.MainChain = "solana"
	req.PrivateKey, _ = library.EncryptByAes(gconv.Bytes(privateKeyHex))
	req.Address = address
	err = service.Address.Add(r.GetCtx(), req)
	if err != nil {
		s.FailJsonExit(r, err.Error())
	}

	// 将新地址注册到Helius Webhook
	go registerAddressToHelius(address)

	result := make(map[string]string)
	result["user_id"] = req.UserId
	result["address"] = req.Address
	s.SusJsonExit(r, result)
}

// Withdraw 提现SOL或SPL Token
func (s *solana) Withdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	if err := r.Parse(&req); err != nil {
		s.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	// 检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "solana")
	if coinInfo == nil {
		s.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	req.MainChain = "solana"
	req.CoinToken = coinInfo.Name
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.Add(r.GetCtx(), req)
	if err != nil {
		s.FailJsonExit(r, err.Error())
	}
	s.SusJsonExit(r)
}

// BalanceOf 查询余额
func (s *solana) BalanceOf(r *ghttp.Request) {
	address := gconv.String(r.GetPost("address"))
	contractAddress := gconv.String(r.GetPost("contract_address"))

	result := make(map[string]string)

	if contractAddress == "" || contractAddress == rpc.SOLNativeMint {
		// 查询SOL余额
		lamports, err := rpc.SolanaClient.GetBalance(address)
		if err != nil {
			s.FailJsonExit(r, "查询SOL余额失败")
		}
		balance := decimal.NewFromInt(int64(lamports)).Div(decimal.NewFromInt(rpc.LamportsPerSOL))
		result["balance"] = balance.String()
	} else {
		// 查询SPL Token余额
		amount, decimals, err := rpc.SolanaClient.GetSPLTokenBalance(address, contractAddress)
		if err != nil {
			s.FailJsonExit(r, "查询SPL Token余额失败")
		}
		result["balance"] = amount
		result["decimals"] = gconv.String(decimals)
	}

	s.SusJsonExit(r, result)
}

// SetAddress 设置监控地址
func (s *solana) SetAddress(r *ghttp.Request) {
	address := gconv.String(r.GetPost("address"))
	count, _ := g.Model("address").Where("main_chain", "solana").Where("address", address).FindCount()
	if count > 0 {
		s.FailJsonExit(r, "地址已存在")
	}
	g.Model("address").Data(g.Map{
		"main_chain": "solana",
		"user_id":    gconv.String(r.GetPost("user_id")),
		"address":    address,
	}).Insert()

	cache := cservice.Cache.New()
	cache.Remove(global.SolanaUserAddressList)

	// 注册到Helius Webhook
	go registerAddressToHelius(address)

	s.SusJsonExit(r)
}

// WebhookReceiver Helius Webhook回调接收接口
func (s *solana) WebhookReceiver(r *ghttp.Request) {
	// 验证Webhook Secret
	webhookSecret := g.Config().GetString("helius.webhook_secret")
	if webhookSecret != "" {
		authHeader := r.GetHeader("Authorization")
		if authHeader != webhookSecret {
			g.Log().File("solana-webhook.{Y-m-d}.log").Printf("Webhook签名验证失败，header: %v", authHeader)
			r.Response.WriteStatusExit(401)
			return
		}
	}

	// 读取请求体
	body := r.GetBody()
	if len(body) == 0 {
		r.Response.WriteStatusExit(200)
		return
	}

	g.Log().File("solana-webhook.{Y-m-d}.log").Printf("收到Helius回调，长度: %d", len(body))

	// 解析Helius Enhanced Webhook载荷
	var payload rpc.HeliusWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		g.Log().File("solana-webhook.{Y-m-d}.log").Printf("解析Webhook载荷失败: %v", err)
		r.Response.WriteStatusExit(200)
		return
	}

	ctx := gctx.New()
	userAddresses, _ := service.Address.GetSolanaAllAddress(ctx)
	coinAddressMap, _ := service.Currency.GetSolanaCoinAddress(ctx)
	contractAddress := g.Config().GetString("solana.contract_address")
	addressRecharge := g.Config().GetBool("solana.address_recharge")
	contractRecharge := g.Config().GetBool("solana.contract_recharge")

	userAddrMap := make(map[string]bool)
	for _, addr := range userAddresses {
		userAddrMap[addr] = true
	}
	// 构建币种Mint地址集合（用于过滤SPL Token）
	coinMintSet := make(map[string]bool)
	for mint := range coinAddressMap {
		coinMintSet[mint] = true
	}
	// 构建系统内部地址排除集合（手续费地址、提现地址、归集地址）
	excludeFromAddrs := make(map[string]bool)
	if feeAddr, _ := cservice.SysConfig.GetConfigByKey("sys.solanaFeeAddress"); feeAddr != nil && feeAddr.ConfigValue != "" {
		excludeFromAddrs[feeAddr.ConfigValue] = true
	}
	if withdrawAddr, _ := cservice.SysConfig.GetConfigByKey("sys.solanaWithdrawAddress"); withdrawAddr != nil && withdrawAddr.ConfigValue != "" {
		excludeFromAddrs[withdrawAddr.ConfigValue] = true
	}
	if mergeAddr, _ := cservice.SysConfig.GetConfigByKey("sys.solanaMergeAddress"); mergeAddr != nil && mergeAddr.ConfigValue != "" {
		excludeFromAddrs[mergeAddr.ConfigValue] = true
	}

	queueExchange := &amqp.QueueExchange{
		QuName: _const.SolanaSweepQuName,
		RtKey:  _const.SolanaSweepRtKey,
		ExName: _const.SolanaSweepExName,
		ExType: _const.SolanaSweepExType,
	}
	mq := amqp.New(queueExchange)
	hasMessage := false

	for _, tx := range payload {
		// 合约充值: 检查feePayer或相关账户是否涉及充值合约
		if contractRecharge && contractAddress != "" {
			for _, acc := range tx.AccountData {
				if acc.Account == contractAddress {
					if processHeliusTransaction(mq, &tx, userAddrMap, coinMintSet, excludeFromAddrs, true) {
						hasMessage = true
					}
					break
				}
			}
		}

		// 地址充值: 处理转入用户地址的交易
		if addressRecharge {
			if processHeliusTransaction(mq, &tx, userAddrMap, coinMintSet, excludeFromAddrs, false) {
				hasMessage = true
			}
		}
	}

	if hasMessage {
		mq.Start()
	}

	r.Response.WriteStatusExit(200)
}

// processHeliusTransaction 处理Helius增强交易数据，将匹配的交易投入MQ
func processHeliusTransaction(mq *amqp.RabbitMQ, tx *rpc.HeliusEnhancedTransaction, userAddrMap map[string]bool, coinMintSet map[string]bool, excludeFromAddrs map[string]bool, isContractRecharge bool) bool {
	produced := false

	// 处理SOL原生转账
	for _, nt := range tx.NativeTransfers {
		if !userAddrMap[nt.ToUserAccount] {
			continue
		}
		// 排除系统内部转账（手续费地址、提现地址、归集地址转入的不算充值）
		if excludeFromAddrs[nt.FromUserAccount] {
			continue
		}
		// 跳过过小的转账（可能是手续费补充）
		if nt.Amount < 5000 {
			continue
		}
		solTx := rpc.SolanaTransaction{
			Signature:       tx.Signature,
			FromAddress:     nt.FromUserAccount,
			ToAddress:       nt.ToUserAccount,
			Amount:          gconv.String(nt.Amount),
			Mint:            "",
			IsToken:         false,
			Slot:            tx.Slot,
			BlockHash:       "",
			Timestamp:       tx.Timestamp,
			TransactionType: tx.Type,
		}
		dd, _ := json.Marshal(solTx)
		mq.RegisterProducer(&producer.SolanaProducer{Msg: string(dd)})
		g.Log().File("solana-producer.{Y-m-d}.log").Printf("Webhook生产 SOL转账 %v → %v, sig: %v", nt.FromUserAccount, nt.ToUserAccount, tx.Signature)
		produced = true
	}

	// 处理SPL Token转账
	for _, tt := range tx.TokenTransfers {
		if !userAddrMap[tt.ToUserAccount] {
			continue
		}
		// 排除系统内部转账
		if excludeFromAddrs[tt.FromUserAccount] {
			continue
		}
		// 检查Token Mint是否在我们的币种列表中
		if !coinMintSet[tt.Mint] {
			continue
		}
		// 计算金额字符串
		amount := decimal.NewFromFloat(tt.TokenAmount)
		solTx := rpc.SolanaTransaction{
			Signature:       tx.Signature,
			FromAddress:     tt.FromUserAccount,
			ToAddress:       tt.ToUserAccount,
			Amount:          amount.String(),
			Mint:            tt.Mint,
			IsToken:         true,
			Slot:            tx.Slot,
			BlockHash:       "",
			Timestamp:       tx.Timestamp,
			TransactionType: tx.Type,
		}
		dd, _ := json.Marshal(solTx)
		mq.RegisterProducer(&producer.SolanaProducer{Msg: string(dd)})
		g.Log().File("solana-producer.{Y-m-d}.log").Printf("Webhook生产 SPL Token转账 mint=%v, %v → %v, sig: %v", tt.Mint, tt.FromUserAccount, tt.ToUserAccount, tx.Signature)
		produced = true
	}

	return produced
}

// registerAddressToHelius 将地址注册到Helius Webhook
func registerAddressToHelius(address string) {
	heliusApiKey := g.Config().GetString("helius.api_key")
	webhookID := g.Config().GetString("helius.webhook_id")
	if heliusApiKey == "" || webhookID == "" {
		g.Log().File("solana-webhook.{Y-m-d}.log").Printf("Helius未配置，跳过注册地址: %v", address)
		return
	}

	client := rpc.NewHeliusClient()
	err := client.AppendAddressesToWebhook(webhookID, []string{address})
	if err != nil {
		g.Log().File("solana-webhook.{Y-m-d}.log").Printf("注册地址到Helius失败: %v, 地址: %v", err, address)
		return
	}
	g.Log().File("solana-webhook.{Y-m-d}.log").Printf("成功注册地址到Helius: %v", address)
}
