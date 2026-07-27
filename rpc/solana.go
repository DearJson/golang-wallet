package rpc

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"gfast/app/common/service"
	"gfast/hdwallet"
	"hash"
	"io/ioutil"
	"net/http"
	"strconv"

	"filippo.io/edwards25519"
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

	SPLTokenAccountSize       = 165
	SolanaTxFeeBufferLamports = 10000
	maxSolanaTransferOutputs  = 2
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
type SolanaRpcClient struct {
	rpcCallOverride func(method string, params interface{}) (json.RawMessage, error)
}

var SolanaClient = &SolanaRpcClient{}

// SolanaRPCError means the RPC server received the request and explicitly
// rejected it. This is different from a transport error where the caller
// cannot know whether the server accepted the transaction.
type SolanaRPCError struct {
	Code    int
	Message string
	Data    json.RawMessage
}

func (e *SolanaRPCError) Error() string {
	return fmt.Sprintf("rpc error [%d]: %s%s", e.Code, e.Message, formatSolanaRpcErrorData(e.Data))
}

func IsSolanaRPCError(err error) bool {
	var rpcErr *SolanaRPCError
	return errors.As(err, &rpcErr)
}

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
	if c.rpcCallOverride != nil {
		return c.rpcCallOverride(method, params)
	}
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
		return nil, &SolanaRPCError{
			Code:    rpcResp.Error.Code,
			Message: rpcResp.Error.Message,
			Data:    rpcResp.Error.Data,
		}
	}
	return rpcResp.Result, nil
}

func formatSolanaRpcErrorData(data json.RawMessage) string {
	if len(data) == 0 || string(data) == "null" {
		return ""
	}
	var detail struct {
		Logs []string    `json:"logs"`
		Err  interface{} `json:"err"`
	}
	if err := json.Unmarshal(data, &detail); err == nil && (len(detail.Logs) > 0 || detail.Err != nil) {
		msg := ""
		if detail.Err != nil {
			if errBytes, err := json.Marshal(detail.Err); err == nil {
				msg += fmt.Sprintf(", data.err=%s", string(errBytes))
			}
		}
		if len(detail.Logs) > 0 {
			logBytes, _ := json.Marshal(detail.Logs)
			msg += fmt.Sprintf(", logs=%s", string(logBytes))
		}
		return msg
	}
	return fmt.Sprintf(", data=%s", string(data))
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

func (c *SolanaRpcClient) getAccountInfo(address string) (*solanaAccountInfo, error) {
	params := []interface{}{
		address,
		map[string]string{"encoding": "jsonParsed"},
	}
	result, err := c.rpcCall("getAccountInfo", params)
	if err != nil {
		return nil, err
	}
	var accountResult struct {
		Value *solanaAccountInfo `json:"value"`
	}
	if err = json.Unmarshal(result, &accountResult); err != nil {
		return nil, err
	}
	return accountResult.Value, nil
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

// GetBlockHeight 获取当前区块高度。
func (c *SolanaRpcClient) GetBlockHeight() (uint64, error) {
	result, err := c.rpcCall("getBlockHeight", []interface{}{})
	if err != nil {
		return 0, err
	}
	var height uint64
	if err = json.Unmarshal(result, &height); err != nil {
		return 0, err
	}
	return height, nil
}

func (c *SolanaRpcClient) GetMinimumBalanceForRentExemption(dataSize uint64) (uint64, error) {
	result, err := c.rpcCall("getMinimumBalanceForRentExemption", []interface{}{dataSize})
	if err != nil {
		return 0, err
	}
	var lamports uint64
	if err = json.Unmarshal(result, &lamports); err != nil {
		return 0, err
	}
	return lamports, nil
}

// ==================== 写操作（构建+签名+发送交易） ====================

// PreparedSolanaTransaction 是已签名、尚未广播的 Solana 交易。
type PreparedSolanaTransaction struct {
	Signature            string
	RawTransaction       []byte
	LastValidBlockHeight uint64
}

// SolanaTransfer 是同一币种的一笔转账输出，Amount 使用币种最小单位。
type SolanaTransfer struct {
	ToAddress string
	Amount    uint64
}

func newPreparedSolanaTransaction(signedTx []byte, lastValidBlockHeight uint64) (*PreparedSolanaTransaction, error) {
	// 当前交易固定只有一个签名，序列化格式为：签名数(1 byte) + 64 byte签名 + message。
	if len(signedTx) < 65 || signedTx[0] != 1 {
		return nil, fmt.Errorf("invalid signed transaction")
	}
	return &PreparedSolanaTransaction{
		Signature:            hdwallet.SolanaBase58Encode(signedTx[1:65]),
		RawTransaction:       signedTx,
		LastValidBlockHeight: lastValidBlockHeight,
	}, nil
}

// PrepareSOL 构建并签名 SOL 转账，但不广播。
func (c *SolanaRpcClient) PrepareSOL(privateKeyHex string, toAddress string, lamports uint64) (*PreparedSolanaTransaction, error) {
	return c.PrepareSOLTransfers(privateKeyHex, []SolanaTransfer{{ToAddress: toAddress, Amount: lamports}})
}

// PrepareSOLTransfers 构建并签名包含多个 SOL 输出的原子交易，但不广播。
func (c *SolanaRpcClient) PrepareSOLTransfers(privateKeyHex string, transfers []SolanaTransfer) (*PreparedSolanaTransaction, error) {
	if len(transfers) == 0 || len(transfers) > maxSolanaTransferOutputs {
		return nil, fmt.Errorf("SOL transfer output count must be between 1 and %d", maxSolanaTransferOutputs)
	}

	// 1. 解析私钥
	privKey, err := hdwallet.GetSolanaPrivateKey(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("get private key error: %v", err)
	}
	fromPubKey := privKey.Public().(ed25519.PublicKey)
	fromAddress := hdwallet.SolanaBase58Encode(fromPubKey)

	// 2. 解析目标地址并计算总额
	toPubKeys := make([][]byte, 0, len(transfers))
	var totalLamports uint64
	seenAddresses := make(map[string]struct{}, len(transfers))
	for _, transfer := range transfers {
		if transfer.Amount == 0 {
			return nil, fmt.Errorf("SOL transfer amount must be greater than zero: address=%s", transfer.ToAddress)
		}
		if _, exists := seenAddresses[transfer.ToAddress]; exists {
			return nil, fmt.Errorf("duplicate SOL transfer address: %s", transfer.ToAddress)
		}
		seenAddresses[transfer.ToAddress] = struct{}{}
		toPubKey, decodeErr := hdwallet.SolanaBase58Decode(transfer.ToAddress)
		if decodeErr != nil || len(toPubKey) != ed25519.PublicKeySize {
			return nil, fmt.Errorf("decode to address error: address=%s err=%v", transfer.ToAddress, decodeErr)
		}
		if bytes.Equal(fromPubKey, toPubKey) {
			return nil, fmt.Errorf("SOL source and destination address are the same: %s", transfer.ToAddress)
		}
		if transfer.Amount > ^uint64(0)-totalLamports {
			return nil, fmt.Errorf("SOL transfer amount overflow")
		}
		totalLamports += transfer.Amount
		toPubKeys = append(toPubKeys, toPubKey)
	}
	fromSOLBalance, err := c.GetBalance(fromAddress)
	if err != nil {
		return nil, fmt.Errorf("get source SOL balance error: %v", err)
	}
	if totalLamports > ^uint64(0)-SolanaTxFeeBufferLamports || fromSOLBalance < totalLamports+SolanaTxFeeBufferLamports {
		return nil, fmt.Errorf(
			"insufficient SOL balance: owner=%s balance=%d transfer=%d fee_buffer=%d",
			fromAddress,
			fromSOLBalance,
			totalLamports,
			SolanaTxFeeBufferLamports,
		)
	}

	// 3. 解析System Program ID
	systemProgramID, _ := hdwallet.SolanaBase58Decode(SystemProgramIDBase58)

	// 4. 获取最新blockhash
	bhResult, err := c.GetLatestBlockhash()
	if err != nil {
		return nil, fmt.Errorf("get blockhash error: %v", err)
	}
	recentBlockhash, err := hdwallet.SolanaBase58Decode(bhResult.Value.Blockhash)
	if err != nil {
		return nil, fmt.Errorf("decode blockhash error: %v", err)
	}

	// 5. 每个输出构建一条 System Program Transfer 指令。
	accountKeys := make([][]byte, 0, len(toPubKeys)+2)
	accountKeys = append(accountKeys, fromPubKey)
	accountKeys = append(accountKeys, toPubKeys...)
	accountKeys = append(accountKeys, systemProgramID)
	programIDIndex := byte(len(accountKeys) - 1)
	instructions := make([]solanaInstruction, 0, len(transfers))
	for i, transfer := range transfers {
		instructionData := make([]byte, 12)
		binary.LittleEndian.PutUint32(instructionData[0:4], 2)
		binary.LittleEndian.PutUint64(instructionData[4:12], transfer.Amount)
		instructions = append(instructions, solanaInstruction{
			programIDIndex: programIDIndex,
			accounts:       []byte{0, byte(i + 1)},
			data:           instructionData,
		})
	}

	// 6. 构建交易消息
	tx := buildTransaction(
		[]ed25519.PublicKey{fromPubKey}, // signers
		accountKeys,
		recentBlockhash,
		instructions,
		1, // numRequiredSignatures
		0, // numReadonlySignedAccounts
		1, // numReadonlyUnsignedAccounts (system program)
	)

	// 7. 签名
	signedTx := signSolanaTransaction(tx, privKey)

	return newPreparedSolanaTransaction(signedTx, bhResult.Value.LastValidBlockHeight)
}

// TransferSOL 转账SOL。
func (c *SolanaRpcClient) TransferSOL(privateKeyHex string, toAddress string, lamports uint64) (string, error) {
	prepared, err := c.PrepareSOL(privateKeyHex, toAddress, lamports)
	if err != nil {
		return "", err
	}
	return c.SendPreparedTransaction(prepared)
}

// PrepareSPLToken 构建并签名 SPL Token 转账，但不广播。
// privateKeyHex: 私钥hex字符串
// toAddress: 目标地址(Base58)
// mint: Token Mint地址(Base58)
// amount: 转账金额（最小单位）
// 返回: 交易签名, error
func (c *SolanaRpcClient) PrepareSPLToken(privateKeyHex string, toAddress string, mint string, amount uint64) (*PreparedSolanaTransaction, error) {
	return c.PrepareSPLTokenTransfers(privateKeyHex, mint, []SolanaTransfer{{ToAddress: toAddress, Amount: amount}})
}

// PrepareSPLTokenTransfers 构建并签名包含多个同 Mint SPL Token 输出的原子交易，但不广播。
func (c *SolanaRpcClient) PrepareSPLTokenTransfers(privateKeyHex string, mint string, transfers []SolanaTransfer) (*PreparedSolanaTransaction, error) {
	if len(transfers) == 0 || len(transfers) > maxSolanaTransferOutputs {
		return nil, fmt.Errorf("SPL token transfer output count must be between 1 and %d", maxSolanaTransferOutputs)
	}

	// 1. 解析私钥
	privKey, err := hdwallet.GetSolanaPrivateKey(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("get private key error: %v", err)
	}
	fromPubKey := privKey.Public().(ed25519.PublicKey)
	fromAddress := hdwallet.SolanaBase58Encode(fromPubKey)

	// 2. 解析各种地址
	mintPubKey, err := hdwallet.SolanaBase58Decode(mint)
	if err != nil || len(mintPubKey) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("decode mint error: %v", err)
	}
	tokenProgramID, _ := hdwallet.SolanaBase58Decode(TokenProgramIDBase58)
	assocTokenProgramID, _ := hdwallet.SolanaBase58Decode(AssociatedTokenProgramIDBase58)
	systemProgramID, _ := hdwallet.SolanaBase58Decode(SystemProgramIDBase58)

	// 3. 校验来源 ATA 与余额。
	fromATA := findAssociatedTokenAddress(fromPubKey, mintPubKey)
	if len(fromATA) != 32 {
		return nil, fmt.Errorf("derive associated token account failed")
	}
	fromATAAddr := hdwallet.SolanaBase58Encode(fromATA)
	fromTokenAccount, err := c.getSPLTokenAccountInfo(fromATAAddr)
	if err != nil {
		return nil, fmt.Errorf("get source token account error: %v", err)
	}
	if fromTokenAccount == nil {
		return nil, fmt.Errorf("source token account not found: owner=%s mint=%s ata=%s", hdwallet.SolanaBase58Encode(fromPubKey), mint, fromATAAddr)
	}
	if fromTokenAccount.Mint != mint {
		return nil, fmt.Errorf("source token account mint mismatch: ata=%s expected=%s actual=%s", fromATAAddr, mint, fromTokenAccount.Mint)
	}
	if fromTokenAccount.Owner != hdwallet.SolanaBase58Encode(fromPubKey) {
		return nil, fmt.Errorf("source token account owner mismatch: ata=%s expected=%s actual=%s", fromATAAddr, hdwallet.SolanaBase58Encode(fromPubKey), fromTokenAccount.Owner)
	}

	type destination struct {
		ownerPubKey []byte
		ata         []byte
		account     *splTokenAccountInfo
		amount      uint64
	}
	destinations := make([]destination, 0, len(transfers))
	seenAddresses := make(map[string]struct{}, len(transfers))
	var totalAmount uint64
	for _, transfer := range transfers {
		if transfer.Amount == 0 {
			return nil, fmt.Errorf("SPL token transfer amount must be greater than zero: address=%s", transfer.ToAddress)
		}
		if _, exists := seenAddresses[transfer.ToAddress]; exists {
			return nil, fmt.Errorf("duplicate SPL token transfer address: %s", transfer.ToAddress)
		}
		seenAddresses[transfer.ToAddress] = struct{}{}
		ownerPubKey, decodeErr := hdwallet.SolanaBase58Decode(transfer.ToAddress)
		if decodeErr != nil || len(ownerPubKey) != ed25519.PublicKeySize {
			return nil, fmt.Errorf("decode to address error: address=%s err=%v", transfer.ToAddress, decodeErr)
		}
		toATA := findAssociatedTokenAddress(ownerPubKey, mintPubKey)
		if len(toATA) != 32 {
			return nil, fmt.Errorf("derive destination associated token account failed: owner=%s", transfer.ToAddress)
		}
		if bytes.Equal(fromATA, toATA) {
			return nil, fmt.Errorf("invalid transfer: source and destination token account are the same: owner=%s mint=%s", transfer.ToAddress, mint)
		}
		if transfer.Amount > ^uint64(0)-totalAmount {
			return nil, fmt.Errorf("SPL token transfer amount overflow")
		}
		totalAmount += transfer.Amount
		toATAAddr := hdwallet.SolanaBase58Encode(toATA)
		toTokenAccount, accountErr := c.getSPLTokenAccountInfo(toATAAddr)
		if accountErr != nil {
			return nil, fmt.Errorf("get destination token account error: owner=%s err=%v", transfer.ToAddress, accountErr)
		}
		if toTokenAccount != nil {
			if toTokenAccount.Mint != mint {
				return nil, fmt.Errorf("destination token account mint mismatch: ata=%s expected=%s actual=%s", toATAAddr, mint, toTokenAccount.Mint)
			}
			if toTokenAccount.Owner != transfer.ToAddress {
				return nil, fmt.Errorf("destination token account owner mismatch: ata=%s expected=%s actual=%s", toATAAddr, transfer.ToAddress, toTokenAccount.Owner)
			}
		}
		destinations = append(destinations, destination{ownerPubKey: ownerPubKey, ata: toATA, account: toTokenAccount, amount: transfer.Amount})
	}
	if fromTokenAccount.Amount < totalAmount {
		return nil, fmt.Errorf("insufficient SPL token balance: owner=%s mint=%s ata=%s balance=%d amount=%d", hdwallet.SolanaBase58Encode(fromPubKey), mint, fromATAAddr, fromTokenAccount.Amount, totalAmount)
	}

	fromSOLBalance, err := c.GetBalance(fromAddress)
	if err != nil {
		return nil, fmt.Errorf("get source SOL balance error: %v", err)
	}
	requiredLamports := uint64(SolanaTxFeeBufferLamports)
	missingATACount := 0
	for _, destination := range destinations {
		if destination.account == nil {
			missingATACount++
		}
	}
	if missingATACount > 0 {
		ataRentLamports, err := c.GetMinimumBalanceForRentExemption(SPLTokenAccountSize)
		if err != nil {
			return nil, fmt.Errorf("get token account rent exemption error: %v", err)
		}
		if ataRentLamports > 0 && uint64(missingATACount) > (^uint64(0)-requiredLamports)/ataRentLamports {
			return nil, fmt.Errorf("required SOL balance overflow")
		}
		requiredLamports += ataRentLamports * uint64(missingATACount)
	}
	if fromSOLBalance < requiredLamports {
		return nil, fmt.Errorf(
			"insufficient SOL for SPL transfer: owner=%s balance=%d required=%d fee_buffer=%d create_destination_ata=%t destination_owner=%s destination_ata=%s mint=%s",
			fromAddress,
			fromSOLBalance,
			requiredLamports,
			SolanaTxFeeBufferLamports,
			missingATACount > 0,
			transfers[0].ToAddress,
			hdwallet.SolanaBase58Encode(destinations[0].ata),
			mint,
		)
	}

	// 4. 获取最新blockhash
	bhResult, err := c.GetLatestBlockhash()
	if err != nil {
		return nil, fmt.Errorf("get blockhash error: %v", err)
	}
	recentBlockhash, err := hdwallet.SolanaBase58Decode(bhResult.Value.Blockhash)
	if err != nil {
		return nil, fmt.Errorf("decode blockhash error: %v", err)
	}

	// 5. 目标 ATA 不存在时先创建，然后逐一转账。
	var instructions []solanaInstruction
	var accountKeys [][]byte

	if missingATACount > 0 {
		sysvarRent, _ := hdwallet.SolanaBase58Decode(SysvarRentPubkeyBase58)
		accountKeys = append(accountKeys, fromPubKey, fromATA)
		for _, destination := range destinations {
			accountKeys = append(accountKeys, destination.ata)
		}
		ownerStart := len(accountKeys)
		for _, destination := range destinations {
			accountKeys = append(accountKeys, destination.ownerPubKey)
		}
		mintIndex := len(accountKeys)
		accountKeys = append(accountKeys, mintPubKey)
		systemIndex := len(accountKeys)
		accountKeys = append(accountKeys, systemProgramID)
		tokenIndex := len(accountKeys)
		accountKeys = append(accountKeys, tokenProgramID)
		rentIndex := len(accountKeys)
		accountKeys = append(accountKeys, sysvarRent)
		associatedTokenIndex := len(accountKeys)
		accountKeys = append(accountKeys, assocTokenProgramID)

		for i, destination := range destinations {
			destinationATAIndex := byte(i + 2)
			if destination.account == nil {
				instructions = append(instructions, solanaInstruction{
					programIDIndex: byte(associatedTokenIndex),
					accounts:       []byte{0, destinationATAIndex, byte(ownerStart + i), byte(mintIndex), byte(systemIndex), byte(tokenIndex), byte(rentIndex)},
					data:           []byte{},
				})
			}
			transferData := make([]byte, 9)
			transferData[0] = 3
			binary.LittleEndian.PutUint64(transferData[1:9], destination.amount)
			instructions = append(instructions, solanaInstruction{
				programIDIndex: byte(tokenIndex),
				accounts:       []byte{1, destinationATAIndex, 0},
				data:           transferData,
			})
		}

		readonlyUnsigned := byte(len(destinations) + 5)
		signedTx := buildAndSignTransaction(privKey, accountKeys, recentBlockhash, instructions, 1, 0, readonlyUnsigned)
		return newPreparedSolanaTransaction(signedTx, bhResult.Value.LastValidBlockHeight)
	}

	accountKeys = append(accountKeys, fromPubKey, fromATA)
	for _, destination := range destinations {
		accountKeys = append(accountKeys, destination.ata)
	}
	tokenIndex := byte(len(accountKeys))
	accountKeys = append(accountKeys, tokenProgramID)
	for i, destination := range destinations {
		transferData := make([]byte, 9)
		transferData[0] = 3
		binary.LittleEndian.PutUint64(transferData[1:9], destination.amount)
		instructions = append(instructions, solanaInstruction{
			programIDIndex: tokenIndex,
			accounts:       []byte{1, byte(i + 2), 0},
			data:           transferData,
		})
	}

	signedTx := buildAndSignTransaction(privKey, accountKeys, recentBlockhash, instructions, 1, 0, 1)
	return newPreparedSolanaTransaction(signedTx, bhResult.Value.LastValidBlockHeight)
}

// TransferSPLToken 转账SPL Token。
func (c *SolanaRpcClient) TransferSPLToken(privateKeyHex string, toAddress string, mint string, amount uint64) (string, error) {
	prepared, err := c.PrepareSPLToken(privateKeyHex, toAddress, mint, amount)
	if err != nil {
		return "", err
	}
	return c.SendPreparedTransaction(prepared)
}

// sendTransaction 发送已签名的交易
func (c *SolanaRpcClient) sendTransaction(signedTx []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(signedTx)
	params := []interface{}{encoded, map[string]interface{}{
		"encoding":            "base64",
		"preflightCommitment": "confirmed",
		"maxRetries":          5,
	}}
	result, err := c.rpcCall("sendTransaction", params)
	if err != nil {
		return "", fmt.Errorf("send transaction error: %w", err)
	}
	var txSig string
	if err = json.Unmarshal(result, &txSig); err != nil {
		return "", fmt.Errorf("unmarshal tx signature error: %v", err)
	}
	return txSig, nil
}

// SendPreparedTransaction 广播已签名的交易。即使 RPC 响应丢失，调用方仍可用预先得到的 Signature 查询交易。
func (c *SolanaRpcClient) SendPreparedTransaction(prepared *PreparedSolanaTransaction) (string, error) {
	if prepared == nil || prepared.Signature == "" || len(prepared.RawTransaction) == 0 {
		return "", fmt.Errorf("invalid prepared transaction")
	}
	signature, err := c.sendTransaction(prepared.RawTransaction)
	if err != nil {
		return prepared.Signature, err
	}
	if signature != prepared.Signature {
		return prepared.Signature, fmt.Errorf("rpc returned unexpected signature: got %s want %s", signature, prepared.Signature)
	}
	return signature, nil
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

type splTokenAccountInfo struct {
	Mint     string
	Owner    string
	Amount   uint64
	Decimals uint8
}

type solanaAccountInfo struct {
	Data struct {
		Parsed struct {
			Info struct {
				Mint        string `json:"mint"`
				Owner       string `json:"owner"`
				TokenAmount struct {
					Amount   string `json:"amount"`
					Decimals uint8  `json:"decimals"`
				} `json:"tokenAmount"`
			} `json:"info"`
			Type string `json:"type"`
		} `json:"parsed"`
		Program string `json:"program"`
	} `json:"data"`
	Owner string `json:"owner"`
}

func (c *SolanaRpcClient) getSPLTokenAccountInfo(address string) (*splTokenAccountInfo, error) {
	account, err := c.getAccountInfo(address)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}
	if account.Owner != TokenProgramIDBase58 || account.Data.Program != "spl-token" || account.Data.Parsed.Type != "account" {
		return nil, fmt.Errorf("account %s is not an SPL token account", address)
	}
	amount, err := strconv.ParseUint(account.Data.Parsed.Info.TokenAmount.Amount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse token balance error: %v", err)
	}
	return &splTokenAccountInfo{
		Mint:     account.Data.Parsed.Info.Mint,
		Owner:    account.Data.Parsed.Info.Owner,
		Amount:   amount,
		Decimals: account.Data.Parsed.Info.TokenAmount.Decimals,
	}, nil
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
// 使用 edwards25519 点解码做严格判断，避免误判导致PDA推导错误。
func isOnCurve(point []byte) bool {
	if len(point) != 32 {
		return false
	}
	_, err := new(edwards25519.Point).SetBytes(point)
	return err == nil
}
