package api

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
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
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/shopspring/decimal"
)

const (
	solanaDepositInstructionIndex            = byte(1)
	solanaDepositPythiaInstructionIndex      = byte(7)
	solanaDepositV5InstructionIndex          = byte(9)
	solanaDepositUsdtReserveInstructionIndex = byte(10)
	solanaDepositUsdtSplitInstructionIndex   = byte(11)
	solanaDepositPythiaV2InstructionIndex    = byte(12)
	solanaDepositPythiaV2ProgramID           = "HhCGLqMrRgoU4M1ymA42gbtKmkbQCVYBz3oDWZziBWAo"
	solanaPythiaMint                         = "CreiuhfwdWCN5mJbMJtA9bBpYQrQF2tCBuZwSPWfpump"
	solanaPythiaTargetOwner                  = "H9YbVf3czoV8fhzQDsGJSRHA3a5qq3bCLXXfSTBTYTaq"
	solanaUsdtReserveAddress                 = "4dVysUQPoLeeC8W57GqQhM9BpyWWw2QZiWtQPdM96F8R"
	solanaDepositUsdtSplitAuthorization      = "7Pp9D2oAjA59rvuFKVws5fcHr7zzQPb7uWa85x135k9z"
	solanaDepositPythiaV2Authorization       = "7Pp9D2oAjA59rvuFKVws5fcHr7zzQPb7uWa85x135k9z"
)

type solanaContractDepositInstruction struct {
	Index                 byte
	Amount                uint64
	IssuedAt              int64
	StudioWallet          string
	Line0Wallet           string
	OperationCenterWallet string
	OrderId               string
	Recipients            []string
	Amounts               []uint64
}

type solana struct {
	WebBase
}

var Solana = new(solana)

func parseSolanaContractDepositInstruction(ixData []byte) (*solanaContractDepositInstruction, error) {
	if len(ixData) < 9 {
		return nil, errors.New("instruction data too short")
	}

	result := &solanaContractDepositInstruction{
		Index:  ixData[0],
		Amount: binary.LittleEndian.Uint64(ixData[1:9]),
	}
	if result.Amount == 0 {
		return nil, errors.New("deposit amount is zero")
	}

	var (
		orderLenOffset int
		orderOffset    int
	)
	switch result.Index {
	case solanaDepositInstructionIndex:
		orderLenOffset = 41
		orderOffset = 45
	case solanaDepositPythiaInstructionIndex:
		orderLenOffset = 9
		orderOffset = 13
	case solanaDepositV5InstructionIndex:
		orderLenOffset = 105
		orderOffset = 109
	case solanaDepositUsdtReserveInstructionIndex:
		orderLenOffset = 9
		orderOffset = 13
	case solanaDepositUsdtSplitInstructionIndex:
		if err := parseSolanaDepositUsdtSplitInstruction(ixData, result); err != nil {
			return nil, err
		}
		return result, nil
	case solanaDepositPythiaV2InstructionIndex:
		if err := parseSolanaDepositPythiaV2Instruction(ixData, result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, errors.New("unsupported contract deposit instruction")
	}

	if len(ixData) < orderOffset {
		return nil, errors.New("instruction order data too short")
	}
	orderIdLen := binary.LittleEndian.Uint32(ixData[orderLenOffset:orderOffset])
	if len(ixData) < orderOffset+int(orderIdLen) {
		return nil, errors.New("instruction order id truncated")
	}
	result.OrderId = string(ixData[orderOffset : orderOffset+int(orderIdLen)])
	if result.Index == solanaDepositUsdtReserveInstructionIndex && len([]byte(result.OrderId)) > 32 {
		return nil, errors.New("instruction order id too long")
	}

	return result, nil
}

func parseSolanaDepositPythiaV2Instruction(ixData []byte, result *solanaContractDepositInstruction) error {
	offset := 9
	issuedAt, err := readBorshI64(ixData, &offset)
	if err != nil {
		return err
	}
	result.IssuedAt = issuedAt

	if len(ixData) < offset+32*3 {
		return errors.New("pythia v2 wallet pubkey truncated")
	}
	result.StudioWallet = hdwallet.SolanaBase58Encode(ixData[offset : offset+32])
	offset += 32
	result.Line0Wallet = hdwallet.SolanaBase58Encode(ixData[offset : offset+32])
	offset += 32
	result.OperationCenterWallet = hdwallet.SolanaBase58Encode(ixData[offset : offset+32])
	offset += 32

	orderIdLen, err := readBorshU32(ixData, &offset)
	if err != nil {
		return err
	}
	if len(ixData) < offset+int(orderIdLen) {
		return errors.New("pythia v2 order id truncated")
	}
	result.OrderId = string(ixData[offset : offset+int(orderIdLen)])
	if len([]byte(result.OrderId)) > 32 {
		return errors.New("pythia v2 order id too long")
	}
	return nil
}

func parseSolanaDepositUsdtSplitInstruction(ixData []byte, result *solanaContractDepositInstruction) error {
	offset := 9
	recipientsLen, err := readBorshU32(ixData, &offset)
	if err != nil {
		return err
	}
	if recipientsLen == 0 {
		return errors.New("split recipients is empty")
	}

	result.Recipients = make([]string, 0, recipientsLen)
	for i := uint32(0); i < recipientsLen; i++ {
		if len(ixData) < offset+32 {
			return errors.New("split recipient pubkey truncated")
		}
		result.Recipients = append(result.Recipients, hdwallet.SolanaBase58Encode(ixData[offset:offset+32]))
		offset += 32
	}

	amountsLen, err := readBorshU32(ixData, &offset)
	if err != nil {
		return err
	}
	if amountsLen != recipientsLen {
		return errors.New("split recipients and amounts length mismatch")
	}

	result.Amounts = make([]uint64, 0, amountsLen)
	var total uint64
	for i := uint32(0); i < amountsLen; i++ {
		amount, err := readBorshU64(ixData, &offset)
		if err != nil {
			return err
		}
		if amount == 0 {
			return errors.New("split amount is zero")
		}
		if total > ^uint64(0)-amount {
			return errors.New("split amount overflow")
		}
		total += amount
		result.Amounts = append(result.Amounts, amount)
	}
	if total != result.Amount {
		return errors.New("split amount sum mismatch")
	}

	orderIdLen, err := readBorshU32(ixData, &offset)
	if err != nil {
		return err
	}
	if len(ixData) < offset+int(orderIdLen) {
		return errors.New("split order id truncated")
	}
	result.OrderId = string(ixData[offset : offset+int(orderIdLen)])
	return nil
}

func readBorshU32(data []byte, offset *int) (uint32, error) {
	if len(data) < *offset+4 {
		return 0, errors.New("borsh u32 truncated")
	}
	value := binary.LittleEndian.Uint32(data[*offset : *offset+4])
	*offset += 4
	return value, nil
}

func readBorshU64(data []byte, offset *int) (uint64, error) {
	if len(data) < *offset+8 {
		return 0, errors.New("borsh u64 truncated")
	}
	value := binary.LittleEndian.Uint64(data[*offset : *offset+8])
	*offset += 8
	return value, nil
}

func readBorshI64(data []byte, offset *int) (int64, error) {
	if len(data) < *offset+8 {
		return 0, errors.New("borsh i64 truncated")
	}
	value := int64(binary.LittleEndian.Uint64(data[*offset : *offset+8]))
	*offset += 8
	return value, nil
}

func decodeSolanaInstructionData(data string) ([]byte, error) {
	ixData, err := hdwallet.SolanaBase58Decode(data)
	if err == nil {
		return ixData, nil
	}
	ixData, base64Err := base64.StdEncoding.DecodeString(data)
	if base64Err == nil {
		return ixData, nil
	}
	return nil, err
}

func isSolanaContractDepositProgram(programID string, configuredProgramID string) bool {
	return programID == configuredProgramID || programID == solanaDepositPythiaV2ProgramID
}

func hasSolanaContractDepositProgram(configuredProgramID string) bool {
	return configuredProgramID != "" || solanaDepositPythiaV2ProgramID != ""
}

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

	if req.Amount <= coinInfo.MinWithdraw {
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
		// 合约充值: 检查交易中是否包含调用充值合约的 Deposit/DepositV5/DepositPythia 指令
		if contractRecharge && hasSolanaContractDepositProgram(contractAddress) {
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

// processContractDeposit 处理合约充值：识别 USDT Deposit(index=1)、DepositV5(index=9)、PYTHIA DepositPythia(index=7)、DepositUsdtReserve(index=10)、DepositUsdtSplit(index=11) 和 DepositPythiaV2(index=12)
func processContractDeposit(mq *amqp.RabbitMQ, tx *rpc.HeliusEnhancedTransaction, contractAddress string, coinMintSet map[string]bool) bool {
	produced := false

	for _, ix := range tx.Instructions {
		if !isSolanaContractDepositProgram(ix.ProgramId, contractAddress) {
			continue
		}

		ixData, err := decodeSolanaInstructionData(ix.Data)
		if err != nil || len(ixData) == 0 {
			continue
		}

		depositIx, err := parseSolanaContractDepositInstruction(ixData)
		if err != nil {
			continue
		}
		if ix.ProgramId == solanaDepositPythiaV2ProgramID && depositIx.Index != solanaDepositPythiaV2InstructionIndex {
			continue
		}

		if len(ix.Accounts) < 1 {
			continue
		}
		userAddress := ix.Accounts[0]

		mint := ""
		totalAmount := decimal.NewFromInt(0)
		toAddress := ix.ProgramId

		switch depositIx.Index {
		case solanaDepositInstructionIndex, solanaDepositV5InstructionIndex:
			// USDT Deposit/DepositV5 会拆分到多个接收方，累加用户所有转出 tokenTransfers 还原总充值额。
			for _, tt := range tx.TokenTransfers {
				if tt.FromUserAccount == userAddress {
					if mint == "" {
						mint = tt.Mint
					}
					totalAmount = totalAmount.Add(decimal.NewFromFloat(tt.TokenAmount))
				}
			}
		case solanaDepositUsdtReserveInstructionIndex:
			toAddress = solanaUsdtReserveAddress
			for _, tt := range tx.TokenTransfers {
				if tt.FromUserAccount != userAddress || tt.ToUserAccount != solanaUsdtReserveAddress {
					continue
				}
				if mint == "" {
					mint = tt.Mint
				}
				totalAmount = totalAmount.Add(decimal.NewFromFloat(tt.TokenAmount))
			}
		case solanaDepositUsdtSplitInstructionIndex:
			if len(ix.Accounts) < 7+len(depositIx.Recipients) {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositUsdtSplit账户数量不足: recipients=%v sig=%v", len(depositIx.Recipients), tx.Signature)
				continue
			}
			authorizationIndex := 6 + len(depositIx.Recipients)
			if ix.Accounts[authorizationIndex] != solanaDepositUsdtSplitAuthorization {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositUsdtSplit授权地址不匹配: got=%v sig=%v", ix.Accounts[authorizationIndex], tx.Signature)
				continue
			}
			mint, totalAmount, err = extractDepositUsdtSplitTransferAmount(tx, userAddress, ix.Accounts[1], ix.Accounts[2:2+len(depositIx.Recipients)], depositIx)
			if err != nil {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositUsdtSplit tokenTransfer校验失败: err=%v user=%v sig=%v", err, userAddress, tx.Signature)
				continue
			}
		case solanaDepositPythiaInstructionIndex:
			if len(ix.Accounts) < 7 {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythia账户数量不足: sig=%v", tx.Signature)
				continue
			}
			mint = solanaPythiaMint
			if len(coinMintSet) > 0 && !coinMintSet[mint] {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythia mint未配置: %v, sig: %v", mint, tx.Signature)
				continue
			}
			decimals, ok := getSolanaTokenDecimalsFromTransfers(tx, mint)
			if !ok {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythia未找到PYTHIA tokenTransfer: sig=%v", tx.Signature)
				continue
			}
			if !hasMatchingPythiaTransfer(tx, userAddress, ix.Accounts[1], ix.Accounts[2], depositIx.Amount) {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythia tokenTransfer校验失败: user=%v, sig=%v", userAddress, tx.Signature)
				continue
			}
			rawAmount, err := decimal.NewFromString(strconv.FormatUint(depositIx.Amount, 10))
			if err != nil {
				continue
			}
			totalAmount = rawAmount.Div(decimal.New(1, int32(decimals)))
		case solanaDepositPythiaV2InstructionIndex:
			if len(ix.Accounts) != 11 {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythiaV2账户数量不匹配: got=%v sig=%v", len(ix.Accounts), tx.Signature)
				continue
			}
			if ix.Accounts[10] != solanaDepositPythiaV2Authorization {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythiaV2授权地址不匹配: got=%v sig=%v", ix.Accounts[10], tx.Signature)
				continue
			}
			mint, totalAmount, err = extractDepositPythiaV2Amount(tx, depositIx)
			if err != nil {
				g.Log().File("solana-producer.{Y-m-d}.log").Printf("DepositPythiaV2金额解析失败: err=%v user=%v sig=%v", err, userAddress, tx.Signature)
				continue
			}
		}

		if totalAmount.IsZero() {
			g.Log().File("solana-producer.{Y-m-d}.log").Printf("合约Deposit index=%v 未找到用户tokenTransfer: user=%v, sig=%v", depositIx.Index, userAddress, tx.Signature)
			continue
		}

		// 如果配置了币种列表，检查 Mint 是否在列表中
		if mint != "" && len(coinMintSet) > 0 && !coinMintSet[mint] {
			g.Log().File("solana-producer.{Y-m-d}.log").Printf("合约Deposit index=%v mint不在币种列表: %v, sig: %v", depositIx.Index, mint, tx.Signature)
			continue
		}

		// 使用 Helius tokenAmount 累加值（人类可读，与地址入金一致）
		amountStr := totalAmount.String()

		solTx := rpc.SolanaTransaction{
			Signature:       tx.Signature,
			FromAddress:     userAddress,
			ToAddress:       toAddress,
			Amount:          amountStr,
			Mint:            mint,
			IsToken:         true,
			Slot:            tx.Slot,
			BlockHash:       "",
			Timestamp:       tx.Timestamp,
			TransactionType: "CONTRACT_DEPOSIT",
			OrderId:         depositIx.OrderId,
		}
		dd, _ := json.Marshal(solTx)
		mq.RegisterProducer(&producer.SolanaProducer{Msg: string(dd)})
		g.Log().File("solana-producer.{Y-m-d}.log").Printf("Webhook生产 合约Deposit index=%v user=%v, amount=%v, mint=%v, order_id=%v, sig: %v", depositIx.Index, userAddress, amountStr, mint, depositIx.OrderId, tx.Signature)
		produced = true
	}

	return produced
}

func extractDepositUsdtSplitTransferAmount(tx *rpc.HeliusEnhancedTransaction, userAddress string, userTokenAccount string, recipientTokenAccounts []string, depositIx *solanaContractDepositInstruction) (string, decimal.Decimal, error) {
	_, ok := getSplitTokenDecimalsFromTransfers(tx, userTokenAccount, recipientTokenAccounts)
	if !ok {
		return "", decimal.Zero, errors.New("token decimals not found")
	}

	mint := ""
	totalAmount := decimal.Zero
	seenRecipients := make([]bool, len(recipientTokenAccounts))
	for _, tt := range tx.TokenTransfers {
		if tt.FromUserAccount != userAddress {
			continue
		}
		if tt.FromTokenAccount != "" && tt.FromTokenAccount != userTokenAccount {
			continue
		}
		for i, recipientTokenAccount := range recipientTokenAccounts {
			if seenRecipients[i] || tt.ToTokenAccount != recipientTokenAccount {
				continue
			}
			if mint == "" {
				mint = tt.Mint
			} else if mint != tt.Mint {
				return "", decimal.Zero, errors.New("recipient mint mismatch")
			}
			if tt.ToUserAccount != userAddress {
				totalAmount = totalAmount.Add(decimal.NewFromFloat(tt.TokenAmount))
			}
			seenRecipients[i] = true
		}
	}
	for _, seen := range seenRecipients {
		if !seen {
			return "", decimal.Zero, errors.New("recipient transfer missing")
		}
	}
	if totalAmount.IsZero() {
		return "", decimal.Zero, errors.New("recipient transfer amount is zero")
	}

	return mint, totalAmount, nil
}

func extractDepositPythiaV2Amount(tx *rpc.HeliusEnhancedTransaction, depositIx *solanaContractDepositInstruction) (string, decimal.Decimal, error) {
	decimals, ok := getSolanaTokenDecimalsFromTransfers(tx, solanaPythiaMint)
	if !ok {
		return "", decimal.Zero, errors.New("pythia decimals not found")
	}

	rawAmount, err := decimal.NewFromString(strconv.FormatUint(depositIx.Amount, 10))
	if err != nil {
		return "", decimal.Zero, err
	}
	return solanaPythiaMint, rawAmount.Div(decimal.New(1, int32(decimals))), nil
}

func calculatePythiaV2SplitAmounts(amount uint64) (targetAmount uint64, studioAmount uint64, line0Amount uint64, operationCenterAmount uint64) {
	line0Amount = percentOfUint64(amount, 5)
	studioAmount = percentOfUint64(amount, 5)
	operationCenterAmount = percentOfUint64(amount, 10)
	targetAmount = amount - line0Amount - studioAmount - operationCenterAmount
	return
}

func percentOfUint64(amount uint64, percent uint64) uint64 {
	return amount/100*percent + amount%100*percent/100
}

func getSplitTokenDecimalsFromTransfers(tx *rpc.HeliusEnhancedTransaction, userTokenAccount string, recipientTokenAccounts []string) (uint8, bool) {
	accountSet := map[string]bool{userTokenAccount: true}
	for _, account := range recipientTokenAccounts {
		accountSet[account] = true
	}
	for _, accountData := range tx.AccountData {
		for _, balanceChange := range accountData.TokenBalanceChanges {
			if accountSet[balanceChange.TokenAccount] {
				return balanceChange.RawTokenAmount.Decimals, true
			}
		}
	}
	return 0, false
}

func getSolanaTokenDecimalsFromTransfers(tx *rpc.HeliusEnhancedTransaction, mint string) (uint8, bool) {
	for _, accountData := range tx.AccountData {
		for _, balanceChange := range accountData.TokenBalanceChanges {
			if balanceChange.Mint == mint {
				return balanceChange.RawTokenAmount.Decimals, true
			}
		}
	}
	return 0, false
}

func hasMatchingPythiaTransfer(tx *rpc.HeliusEnhancedTransaction, userAddress string, userTokenAccount string, targetTokenAccount string, rawAmount uint64) bool {
	return hasMatchingPythiaTransferToAccount(tx, userAddress, userTokenAccount, targetTokenAccount, solanaPythiaTargetOwner, rawAmount)
}

func hasMatchingPythiaTransferToAccount(tx *rpc.HeliusEnhancedTransaction, userAddress string, userTokenAccount string, targetTokenAccount string, targetOwner string, rawAmount uint64) bool {
	for _, tt := range tx.TokenTransfers {
		if tt.Mint != solanaPythiaMint {
			continue
		}
		if tt.FromUserAccount != userAddress {
			continue
		}
		if tt.FromTokenAccount != "" && tt.FromTokenAccount != userTokenAccount {
			continue
		}
		if tt.ToTokenAccount != targetTokenAccount {
			continue
		}
		if targetOwner != "" && tt.ToUserAccount != "" && tt.ToUserAccount != targetOwner {
			continue
		}
		if !tokenTransferAmountMatches(tx, tt.ToTokenAccount, solanaPythiaMint, rawAmount) {
			continue
		}
		return true
	}
	return false
}

func tokenTransferAmountMatches(tx *rpc.HeliusEnhancedTransaction, tokenAccount string, mint string, rawAmount uint64) bool {
	expectedAmount, err := decimal.NewFromString(strconv.FormatUint(rawAmount, 10))
	if err != nil {
		return false
	}
	for _, accountData := range tx.AccountData {
		for _, balanceChange := range accountData.TokenBalanceChanges {
			if balanceChange.TokenAccount != tokenAccount || balanceChange.Mint != mint {
				continue
			}
			actualAmount, err := decimal.NewFromString(balanceChange.RawTokenAmount.TokenAmount)
			if err != nil {
				continue
			}
			if actualAmount.Abs().Equal(expectedAmount) {
				return true
			}
		}
	}
	return false
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

// ResetHash 通过交易hash手动导入充值记录
// 使用Helius Parse Transaction API获取增强交易数据，复用与Webhook相同的验证逻辑
func (s *solana) ResetHash(r *ghttp.Request) {
	hash := gconv.String(r.GetPost("hash"))
	if hash == "" {
		s.FailJsonExit(r, "交易hash不能为空")
	}

	// 检查交易是否已存在
	ctx := gctx.New()
	exists, _ := service.Recharge.GetInfoByHash(ctx, hash)
	if exists != nil {
		s.FailJsonExit(r, "该交易已存在充值记录")
	}

	// 通过Helius API解析交易，获取与Webhook回调相同格式的增强交易数据
	heliusClient := rpc.NewHeliusClient()
	if heliusClient.ApiKey == "" {
		s.FailJsonExit(r, "Helius API未配置")
	}

	txList, err := heliusClient.ParseTransactions([]string{hash})
	if err != nil {
		g.Log().File("solana-reset.{Y-m-d}.log").Printf("解析交易失败: %v, hash: %v", err, hash)
		s.FailJsonExit(r, "解析交易失败: "+err.Error())
	}
	if len(txList) == 0 {
		s.FailJsonExit(r, "未找到该交易或交易尚未确认")
	}

	// 先通过RPC确认交易状态（必须是成功的交易）
	sigStatus, err := rpc.SolanaClient.GetSignatureStatus(hash)
	if err != nil {
		s.FailJsonExit(r, "查询交易状态失败: "+err.Error())
	}
	if sigStatus == nil {
		s.FailJsonExit(r, "交易不存在或尚未确认")
	}
	if sigStatus.Err != nil {
		s.FailJsonExit(r, "该交易执行失败，无法导入")
	}

	// 加载配置和地址信息（与WebhookReceiver完全相同的逻辑）
	userAddresses, _ := service.Address.GetSolanaAllAddress(ctx)
	coinAddressMap, _ := service.Currency.GetSolanaCoinAddress(ctx)
	contractAddress := g.Config().GetString("solana.contract_address")
	addressRecharge := g.Config().GetBool("solana.address_recharge")
	contractRecharge := g.Config().GetBool("solana.contract_recharge")

	userAddrMap := make(map[string]bool)
	for _, addr := range userAddresses {
		userAddrMap[addr] = true
	}
	coinMintSet := make(map[string]bool)
	for mint := range coinAddressMap {
		coinMintSet[mint] = true
	}
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

	for _, tx := range txList {
		// 合约充值
		if contractRecharge && hasSolanaContractDepositProgram(contractAddress) {
			if processContractDeposit(mq, &tx, contractAddress, coinMintSet) {
				hasMessage = true
			}
		}
		// 地址充值
		if addressRecharge {
			if processHeliusTransaction(mq, &tx, userAddrMap, coinMintSet, excludeFromAddrs, false) {
				hasMessage = true
			}
		}
	}

	if !hasMessage {
		s.FailJsonExit(r, "该交易不满足充值条件（非合约Deposit/DepositV5/DepositPythia/DepositPythiaV2指令、非监控地址或币种不匹配）")
	}

	mq.Start()
	g.Log().File("solana-reset.{Y-m-d}.log").Printf("手动导入交易成功: %v", hash)
	s.SusJsonExit(r)
}
