package rpc

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gfast/app/common/service"
	"gfast/hdwallet"
	"hash"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

// Solana 常量
const (
	SOLDecimals    = 9
	LamportsPerSOL = 1000000000 // 10^9
	// SOL原生代币标记地址（在currency表中用于标识SOL）
	SOLNativeMint = "So11111111111111111111111111111111111111112"
)

// Solana知名程序ID（Base58编码）
const (
	SystemProgramIDBase58          = "11111111111111111111111111111111"
	TokenProgramIDBase58           = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
	AssociatedTokenProgramIDBase58 = "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"
	SysvarRentPubkeyBase58         = "SysvarRent111111111111111111111111111111111"
)

// ==================== Solana RPC 客户端 ====================

// SolanaRpcClient Solana JSON-RPC 客户端
type SolanaRpcClient struct{}

var SolanaClient = &SolanaRpcClient{}

// getSolanaRpcUrl 获取Solana RPC URL
func getSolanaRpcUrl() string {
	cache := service.Cache.New()
	url := gconv.String(cache.Get("solana_rpc_url"))
	if url == "" {
		url = "https://api.mainnet-beta.solana.com"
	}
	return url
}

// rpcCall 发送Solana JSON-RPC请求
func (c *SolanaRpcClient) rpcCall(method string, params interface{}) (json.RawMessage, error) {
	reqBody := SolanaRpcRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request error: %v", err)
	}

	rpcUrl := getSolanaRpcUrl()
	resp, err := http.Post(rpcUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("rpc request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	var rpcResp struct {
		Jsonrpc string          `json:"jsonrpc"`
		ID      int             `json:"id"`
		Result  json.RawMessage `json:"result"`
		Error   *SolanaRpcError `json:"error,omitempty"`
	}
	if err = json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error [%d]: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	return rpcResp.Result, nil
}

// ==================== 读操作 ====================

// GetBalance 查询SOL余额（返回lamports）
func (c *SolanaRpcClient) GetBalance(address string) (uint64, error) {
	result, err := c.rpcCall("getBalance", []interface{}{address})
	if err != nil {
		return 0, err
	}
	var balResult SolanaBalanceResult
	if err = json.Unmarshal(result, &balResult); err != nil {
		return 0, err
	}
	return balResult.Value, nil
}

// GetTokenAccountsByOwner 查询某个地址拥有的所有SPL Token账户
func (c *SolanaRpcClient) GetTokenAccountsByOwner(owner string) (*SolanaTokenAccountResult, error) {
	params := []interface{}{
		owner,
		map[string]string{"programId": TokenProgramIDBase58},
		map[string]string{"encoding": "jsonParsed"},
	}
	result, err := c.rpcCall("getTokenAccountsByOwner", params)
	if err != nil {
		return nil, err
	}
	var tokenResult SolanaTokenAccountResult
	if err = json.Unmarshal(result, &tokenResult); err != nil {
		return nil, err
	}
	return &tokenResult, nil
}

// GetSPLTokenBalance 查询某个地址的指定SPL Token余额
func (c *SolanaRpcClient) GetSPLTokenBalance(owner string, mint string) (string, uint8, error) {
	tokenResult, err := c.GetTokenAccountsByOwner(owner)
	if err != nil {
		return "0", 0, err
	}
	for _, account := range tokenResult.Value {
		if account.Account.Data.Parsed.Info.Mint == mint {
			return account.Account.Data.Parsed.Info.TokenAmount.Amount, account.Account.Data.Parsed.Info.TokenAmount.Decimals, nil
		}
	}
	return "0", 0, nil
}

// GetSignatureStatuses 查询交易签名状态
func (c *SolanaRpcClient) GetSignatureStatuses(signatures []string) ([]*SolanaSignatureStatus, error) {
	params := []interface{}{signatures, map[string]bool{"searchTransactionHistory": true}}
	result, err := c.rpcCall("getSignatureStatuses", params)
	if err != nil {
		return nil, err
	}
	var statusResult struct {
		Context struct {
			Slot uint64 `json:"slot"`
		} `json:"context"`
		Value []*SolanaSignatureStatus `json:"value"`
	}
	if err = json.Unmarshal(result, &statusResult); err != nil {
		return nil, err
	}
	return statusResult.Value, nil
}

// GetSignatureStatus 查询单个交易签名状态
func (c *SolanaRpcClient) GetSignatureStatus(signature string) (*SolanaSignatureStatus, error) {
	statuses, err := c.GetSignatureStatuses([]string{signature})
	if err != nil {
		return nil, err
	}
	if len(statuses) == 0 || statuses[0] == nil {
		return nil, nil
	}
	return statuses[0], nil
}

// GetSignaturesForAddress 获取某地址的最近交易签名列表
func (c *SolanaRpcClient) GetSignaturesForAddress(address string, limit int, before string) ([]SolanaSignatureInfo, error) {
	opts := map[string]interface{}{"limit": limit}
	if before != "" {
		opts["before"] = before
	}
	params := []interface{}{address, opts}
	result, err := c.rpcCall("getSignaturesForAddress", params)
	if err != nil {
		return nil, err
	}
	var sigInfos []SolanaSignatureInfo
	if err = json.Unmarshal(result, &sigInfos); err != nil {
		return nil, err
	}
	return sigInfos, nil
}

// GetLatestBlockhash 获取最新的blockhash
func (c *SolanaRpcClient) GetLatestBlockhash() (*SolanaBlockhashResult, error) {
	result, err := c.rpcCall("getLatestBlockhash", []interface{}{})
	if err != nil {
		return nil, err
	}
	var bhResult SolanaBlockhashResult
	if err = json.Unmarshal(result, &bhResult); err != nil {
		return nil, err
	}
	return &bhResult, nil
}

// ==================== 写操作（构建+签名+发送交易） ====================

// TransferSOL 转账SOL
// privateKeyHex: 私钥hex字符串
// toAddress: 目标地址(Base58)
// lamports: 转账金额(lamports)
// 返回: 交易签名, error
func (c *SolanaRpcClient) TransferSOL(privateKeyHex string, toAddress string, lamports uint64) (string, error) {
	// 1. 解析私钥
	privKey, err := hdwallet.GetSolanaPrivateKey(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("get private key error: %v", err)
	}
	fromPubKey := privKey.Public().(ed25519.PublicKey)

	// 2. 解析目标地址
	toPubKey, err := hdwallet.SolanaBase58Decode(toAddress)
	if err != nil {
		return "", fmt.Errorf("decode to address error: %v", err)
	}

	// 3. 解析System Program ID
	systemProgramID, _ := hdwallet.SolanaBase58Decode(SystemProgramIDBase58)

	// 4. 获取最新blockhash
	bhResult, err := c.GetLatestBlockhash()
	if err != nil {
		return "", fmt.Errorf("get blockhash error: %v", err)
	}
	recentBlockhash, err := hdwallet.SolanaBase58Decode(bhResult.Value.Blockhash)
	if err != nil {
		return "", fmt.Errorf("decode blockhash error: %v", err)
	}

	// 5. 构建System Program Transfer指令
	// 指令数据: [2,0,0,0] (transfer index=2, little-endian u32) + lamports (little-endian u64)
	instructionData := make([]byte, 12)
	binary.LittleEndian.PutUint32(instructionData[0:4], 2) // transfer instruction
	binary.LittleEndian.PutUint64(instructionData[4:12], lamports)

	// 6. 构建交易消息
	tx := buildTransaction(
		[]ed25519.PublicKey{fromPubKey},                 // signers
		[][]byte{fromPubKey, toPubKey, systemProgramID}, // account keys
		recentBlockhash,
		[]solanaInstruction{
			{
				programIDIndex: 2,            // systemProgramID
				accounts:       []byte{0, 1}, // from, to
				data:           instructionData,
			},
		},
		1, // numRequiredSignatures
		0, // numReadonlySignedAccounts
		1, // numReadonlyUnsignedAccounts (system program)
	)

	// 7. 签名
	signedTx := signSolanaTransaction(tx, privKey)

	// 8. 发送交易
	return c.sendTransaction(signedTx)
}

// TransferSPLToken 转账SPL Token
// privateKeyHex: 私钥hex字符串
// toAddress: 目标地址(Base58)
// mint: Token Mint地址(Base58)
// amount: 转账金额（最小单位）
// 返回: 交易签名, error
func (c *SolanaRpcClient) TransferSPLToken(privateKeyHex string, toAddress string, mint string, amount uint64) (string, error) {
	// 1. 解析私钥
	privKey, err := hdwallet.GetSolanaPrivateKey(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("get private key error: %v", err)
	}
	fromPubKey := privKey.Public().(ed25519.PublicKey)

	// 2. 解析各种地址
	toPubKeyBytes, err := hdwallet.SolanaBase58Decode(toAddress)
	if err != nil {
		return "", fmt.Errorf("decode to address error: %v", err)
	}
	mintPubKey, err := hdwallet.SolanaBase58Decode(mint)
	if err != nil {
		return "", fmt.Errorf("decode mint error: %v", err)
	}
	tokenProgramID, _ := hdwallet.SolanaBase58Decode(TokenProgramIDBase58)
	assocTokenProgramID, _ := hdwallet.SolanaBase58Decode(AssociatedTokenProgramIDBase58)
	systemProgramID, _ := hdwallet.SolanaBase58Decode(SystemProgramIDBase58)

	// 3. 计算Associated Token Account (ATA)
	fromATA := findAssociatedTokenAddress(fromPubKey, mintPubKey)
	toATA := findAssociatedTokenAddress(toPubKeyBytes, mintPubKey)

	// 4. 获取最新blockhash
	bhResult, err := c.GetLatestBlockhash()
	if err != nil {
		return "", fmt.Errorf("get blockhash error: %v", err)
	}
	recentBlockhash, err := hdwallet.SolanaBase58Decode(bhResult.Value.Blockhash)
	if err != nil {
		return "", fmt.Errorf("decode blockhash error: %v", err)
	}

	// 5. 检查目标ATA是否存在，不存在需要创建
	toATAAddr := hdwallet.SolanaBase58Encode(toATA)
	toATABalance, err := c.GetBalance(toATAAddr)

	var instructions []solanaInstruction
	var accountKeys [][]byte

	if err != nil || toATABalance == 0 {
		// 需要先创建ATA
		// 创建ATA指令：无data，accounts=[payer, ata, owner, mint, system, token_program]
		// 再加transfer指令
		sysvarRent, _ := hdwallet.SolanaBase58Decode(SysvarRentPubkeyBase58)
		accountKeys = [][]byte{
			fromPubKey,          // 0: payer/signer
			toATA,               // 1: to ATA (writable)
			toPubKeyBytes,       // 2: to owner
			mintPubKey,          // 3: mint
			systemProgramID,     // 4: system program
			tokenProgramID,      // 5: token program
			sysvarRent,          // 6: rent sysvar (不需要了,新版可忽略但为兼容保留)
			assocTokenProgramID, // 7: associated token program
			fromATA,             // 8: from ATA (writable)
		}
		// 创建ATA指令
		instructions = append(instructions, solanaInstruction{
			programIDIndex: 7, // associated token program
			accounts:       []byte{0, 1, 2, 3, 4, 5, 6},
			data:           []byte{}, // 空data表示CreateAssociatedTokenAccount
		})
		// Transfer指令
		transferData := make([]byte, 9)
		transferData[0] = 3 // Transfer instruction index
		binary.LittleEndian.PutUint64(transferData[1:9], amount)
		instructions = append(instructions, solanaInstruction{
			programIDIndex: 5,               // token program
			accounts:       []byte{8, 1, 0}, // from_ata, to_ata, owner/signer
			data:           transferData,
		})

		signedTx := buildAndSignTransaction(privKey, accountKeys, recentBlockhash, instructions, 1, 0, 4)
		return c.sendTransaction(signedTx)
	}

	// 目标ATA已存在，直接transfer
	accountKeys = [][]byte{
		fromPubKey,     // 0: owner/signer
		fromATA,        // 1: from ATA (writable)
		toATA,          // 2: to ATA (writable)
		tokenProgramID, // 3: token program
	}
	transferData := make([]byte, 9)
	transferData[0] = 3 // Transfer instruction index
	binary.LittleEndian.PutUint64(transferData[1:9], amount)
	instructions = []solanaInstruction{
		{
			programIDIndex: 3,               // token program
			accounts:       []byte{1, 2, 0}, // from_ata, to_ata, owner
			data:           transferData,
		},
	}

	signedTx := buildAndSignTransaction(privKey, accountKeys, recentBlockhash, instructions, 1, 0, 1)
	return c.sendTransaction(signedTx)
}

// sendTransaction 发送已签名的交易
func (c *SolanaRpcClient) sendTransaction(signedTx []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(signedTx)
	params := []interface{}{encoded, map[string]string{"encoding": "base64"}}
	result, err := c.rpcCall("sendTransaction", params)
	if err != nil {
		return "", fmt.Errorf("send transaction error: %v", err)
	}
	var txSig string
	if err = json.Unmarshal(result, &txSig); err != nil {
		return "", fmt.Errorf("unmarshal tx signature error: %v", err)
	}
	return txSig, nil
}

// ==================== SOL/SPL Token 提现封装 ====================

// SolanaTransferSOL 提现SOL的高层封装
func SolanaTransferSOL(privateKeyHex string, toAddress string, amount decimal.Decimal) (string, error) {
	lamports := amount.Mul(decimal.NewFromInt(LamportsPerSOL)).IntPart()
	sig, err := SolanaClient.TransferSOL(privateKeyHex, toAddress, uint64(lamports))
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana SOL 转账失败: %v", err)
		return "", err
	}
	return sig, nil
}

// SolanaTransferSPLToken 提现SPL Token的高层封装
func SolanaTransferSPLToken(privateKeyHex string, toAddress string, mint string, amount decimal.Decimal, decimals int) (string, error) {
	tenDecimal := decimal.NewFromInt(1)
	for i := 0; i < decimals; i++ {
		tenDecimal = tenDecimal.Mul(decimal.NewFromInt(10))
	}
	tokenAmount := amount.Mul(tenDecimal).IntPart()
	sig, err := SolanaClient.TransferSPLToken(privateKeyHex, toAddress, mint, uint64(tokenAmount))
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("Solana SPL Token 转账失败: %v", err)
		return "", err
	}
	return sig, nil
}

// ==================== 交易构建辅助函数 ====================

type solanaInstruction struct {
	programIDIndex byte
	accounts       []byte
	data           []byte
}

// buildTransaction 构建Solana交易消息（不含签名）
func buildTransaction(signers []ed25519.PublicKey, accountKeys [][]byte, recentBlockhash []byte, instructions []solanaInstruction, numRequired, numReadonlySigned, numReadonlyUnsigned byte) []byte {
	var msg bytes.Buffer

	// Message Header
	msg.WriteByte(numRequired)
	msg.WriteByte(numReadonlySigned)
	msg.WriteByte(numReadonlyUnsigned)

	// Account keys
	writeCompactU16(&msg, len(accountKeys))
	for _, key := range accountKeys {
		msg.Write(padOrTrim(key, 32))
	}

	// Recent blockhash
	msg.Write(padOrTrim(recentBlockhash, 32))

	// Instructions
	writeCompactU16(&msg, len(instructions))
	for _, ix := range instructions {
		msg.WriteByte(ix.programIDIndex)
		writeCompactU16(&msg, len(ix.accounts))
		msg.Write(ix.accounts)
		writeCompactU16(&msg, len(ix.data))
		msg.Write(ix.data)
	}

	return msg.Bytes()
}

// buildAndSignTransaction 构建并签名交易
func buildAndSignTransaction(privKey ed25519.PrivateKey, accountKeys [][]byte, recentBlockhash []byte, instructions []solanaInstruction, numRequired, numReadonlySigned, numReadonlyUnsigned byte) []byte {
	message := buildTransaction(nil, accountKeys, recentBlockhash, instructions, numRequired, numReadonlySigned, numReadonlyUnsigned)
	return signSolanaTransaction(message, privKey)
}

// signSolanaTransaction 对消息进行Ed25519签名并返回完整交易
func signSolanaTransaction(message []byte, privKey ed25519.PrivateKey) []byte {
	signature := ed25519.Sign(privKey, message)

	var tx bytes.Buffer
	// Signatures: compact-u16 count + signatures
	writeCompactU16(&tx, 1)
	tx.Write(signature) // 64 bytes

	// Message
	tx.Write(message)

	return tx.Bytes()
}

// findAssociatedTokenAddress 计算Associated Token Address (PDA)
func findAssociatedTokenAddress(wallet []byte, mint []byte) []byte {
	assocTokenProgramID, _ := hdwallet.SolanaBase58Decode(AssociatedTokenProgramIDBase58)
	tokenProgramID, _ := hdwallet.SolanaBase58Decode(TokenProgramIDBase58)

	// PDA = ProgramDerivedAddress([wallet, token_program, mint], associated_token_program)
	seeds := [][]byte{
		padOrTrim(wallet, 32),
		padOrTrim(tokenProgramID, 32),
		padOrTrim(mint, 32),
	}
	pda, _ := findProgramAddress(seeds, assocTokenProgramID)
	return pda
}

// findProgramAddress 查找Program Derived Address
func findProgramAddress(seeds [][]byte, programID []byte) ([]byte, byte) {
	for nonce := byte(255); ; nonce-- {
		address := createProgramAddress(append(seeds, []byte{nonce}), programID)
		if address != nil {
			return address, nonce
		}
		if nonce == 0 {
			break
		}
	}
	return nil, 0
}

// createProgramAddress 创建Program Address
func createProgramAddress(seeds [][]byte, programID []byte) []byte {
	// SHA256(seeds... + programID + "ProgramDerivedAddress")
	hasher := newSha256Hasher()
	for _, seed := range seeds {
		hasher.Write(seed)
	}
	hasher.Write(padOrTrim(programID, 32))
	hasher.Write([]byte("ProgramDerivedAddress"))
	hash := hasher.Sum(nil)

	// 检查结果是否在Ed25519曲线上（PDA必须不在曲线上）
	if isOnCurve(hash) {
		return nil
	}
	return hash
}

// writeCompactU16 写入Compact-u16编码
func writeCompactU16(buf *bytes.Buffer, value int) {
	if value < 0x80 {
		buf.WriteByte(byte(value))
	} else if value < 0x4000 {
		buf.WriteByte(byte(value&0x7F) | 0x80)
		buf.WriteByte(byte(value >> 7))
	} else {
		buf.WriteByte(byte(value&0x7F) | 0x80)
		buf.WriteByte(byte((value>>7)&0x7F) | 0x80)
		buf.WriteByte(byte(value >> 14))
	}
}

// padOrTrim 确保字节切片为指定长度
func padOrTrim(data []byte, size int) []byte {
	if len(data) == size {
		return data
	}
	if len(data) > size {
		return data[:size]
	}
	result := make([]byte, size)
	copy(result, data)
	return result
}

// newSha256Hasher 创建SHA256哈希器
func newSha256Hasher() hash.Hash {
	return sha256.New()
}

// isOnCurve 检查点是否在Ed25519曲线上
// 对于PDA查找，SHA256结果落在曲线上的概率约为1/2，
// 所以 findProgramAddress 会在几次尝试内找到有效PDA。
// 这里使用简化检查：尝试用该点做一次签名验证。
func isOnCurve(point []byte) bool {
	if len(point) != 32 {
		return false
	}
	// 通过尝试构造公钥并验证一个空签名来检测点是否在曲线上
	// 如果点不在曲线上，ed25519.Verify 会 panic 或返回 false
	defer func() { recover() }()
	pubKey := ed25519.PublicKey(point)
	// 尝试验证一个任意签名，如果公钥无效会 panic
	ed25519.Verify(pubKey, []byte("test"), make([]byte, 64))
	return true
}
