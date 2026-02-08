package api

import (
	"encoding/binary"
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
		// 合约充值: 检查交易中是否包含调用充值合约的 Deposit 指令（discriminator=1）
		if contractRecharge && contractAddress != "" {
			if processContractDeposit(mq, &tx, contractAddress, coinMintSet) {
				hasMessage = true
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

// processContractDeposit 处理合约充值：仅当交易包含 Deposit 指令（Borsh discriminator=1）时才记录
// 合约指令数据格式: [0x01] + [amount: u64 LE 8字节] + [studio_wallet: Pubkey 32字节]
// 合约账户顺序: [0]user, [1]user_usdt, [2]contract_usdt, [3]studio_usdt, [4]operation_usdt, [5]token_program, [6]state
func processContractDeposit(mq *amqp.RabbitMQ, tx *rpc.HeliusEnhancedTransaction, contractAddress string, coinMintSet map[string]bool) bool {
	produced := false

	for _, ix := range tx.Instructions {
		if ix.ProgramId != contractAddress {
			continue
		}

		// Base58 解码指令数据
		ixData, err := hdwallet.SolanaBase58Decode(ix.Data)
		if err != nil || len(ixData) == 0 {
			continue
		}
		// Deposit 指令 Borsh 布局:
		// [0]     = discriminator 0x01
		// [1-8]   = amount (u64 LE)
		// [9-40]  = studio_wallet (Pubkey, 32 bytes)
		// [41-44] = order_id 长度 (u32 LE)
		// [45+]   = order_id UTF-8 内容
		// 最小长度: 1 + 8 + 32 + 4 = 45 字节（order_id可为空字符串）
		if len(ixData) < 45 || ixData[0] != 1 {
			continue
		}

		// 解析充值金额（u64 小端序）— 仅用于校验非零
		rawAmount := binary.LittleEndian.Uint64(ixData[1:9])
		if rawAmount == 0 {
			continue
		}

		// 解析 order_id（Borsh String = u32长度前缀 + UTF-8内容）
		orderIdLen := binary.LittleEndian.Uint32(ixData[41:45])
		orderId := ""
		if orderIdLen > 0 {
			if len(ixData) < 45+int(orderIdLen) {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("Deposit指令数据不完整，order_id截断: sig=%v", tx.Signature)
				continue
			}
			orderId = string(ixData[45 : 45+orderIdLen])
		}

		// 用户地址 = 指令的第一个账户（accounts[0]）
		if len(ix.Accounts) < 1 {
			continue
		}
		userAddress := ix.Accounts[0]

		// 从 tokenTransfers 中累加用户的总转出金额（Helius tokenAmount 已经是人类可读值，含精度转换）
		// 合约 Deposit 会将用户资金拆分到多个接收方（平台/工作室/运营/合约），需要求和还原总金额
		mint := ""
		totalAmount := decimal.NewFromInt(0)
		for _, tt := range tx.TokenTransfers {
			if tt.FromUserAccount == userAddress {
				if mint == "" {
					mint = tt.Mint
				}
				totalAmount = totalAmount.Add(decimal.NewFromFloat(tt.TokenAmount))
			}
		}

		if totalAmount.IsZero() {
			g.Log().File("solana-producer.{Y-m-d}.log").Printf("合约Deposit未找到用户tokenTransfer: user=%v, sig=%v", userAddress, tx.Signature)
			continue
		}

		// 如果配置了币种列表，检查 Mint 是否在列表中
		if mint != "" && len(coinMintSet) > 0 && !coinMintSet[mint] {
			g.Log().File("solana-producer.{Y-m-d}.log").Printf("合约Deposit mint不在币种列表: %v, sig: %v", mint, tx.Signature)
			continue
		}

		// 使用 Helius tokenAmount 累加值（人类可读，与地址入金一致）
		amountStr := totalAmount.String()

		solTx := rpc.SolanaTransaction{
			Signature:       tx.Signature,
			FromAddress:     userAddress,
			ToAddress:       contractAddress,
			Amount:          amountStr,
			Mint:            mint,
			IsToken:         true,
			Slot:            tx.Slot,
			BlockHash:       "",
			Timestamp:       tx.Timestamp,
			TransactionType: "CONTRACT_DEPOSIT",
			OrderId:         orderId,
		}
		dd, _ := json.Marshal(solTx)
		mq.RegisterProducer(&producer.SolanaProducer{Msg: string(dd)})
		g.Log().File("solana-producer.{Y-m-d}.log").Printf("Webhook生产 合约Deposit user=%v, amount=%v, mint=%v, order_id=%v, sig: %v", userAddress, amountStr, mint, orderId, tx.Signature)
		produced = true
	}

	return produced
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
