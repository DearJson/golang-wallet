package rpc

// ==================== Solana JSON-RPC 请求/响应结构 ====================

// SolanaRpcRequest JSON-RPC 请求
type SolanaRpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// SolanaRpcResponse JSON-RPC 响应
type SolanaRpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  interface{}     `json:"result"`
	Error   *SolanaRpcError `json:"error,omitempty"`
}

// SolanaRpcError JSON-RPC 错误
type SolanaRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ==================== Solana 交易相关结构 ====================

// SolanaTransaction 投入MQ队列的消息结构
type SolanaTransaction struct {
	Signature       string  `json:"signature"`
	FromAddress     string  `json:"from_address"`
	ToAddress       string  `json:"to_address"`
	Amount          string  `json:"amount"`
	Decimals        uint8   `json:"decimals"`
	Mint            string  `json:"mint"`     // SPL Token Mint地址，空表示SOL
	IsToken         bool    `json:"is_token"` // 是否是SPL Token交易
	Slot            uint64  `json:"slot"`
	BlockHash       string  `json:"block_hash"`
	Timestamp       int64   `json:"timestamp"`
	TransactionType string  `json:"transaction_type"` // TRANSFER / UNKNOWN / CONTRACT_DEPOSIT
	Fee             float64 `json:"fee"`
	OrderId         string  `json:"order_id,omitempty"` // 合约充值订单号（从Deposit指令解析）
}

// ==================== Helius Webhook 回调数据结构 ====================

// HeliusWebhookPayload Helius Enhanced Webhook回调载荷（数组）
type HeliusWebhookPayload []HeliusEnhancedTransaction

// HeliusEnhancedTransaction Helius增强交易数据
type HeliusEnhancedTransaction struct {
	AccountData     []HeliusAccountData    `json:"accountData"`
	Description     string                 `json:"description"`
	Events          map[string]interface{} `json:"events"`
	Fee             int64                  `json:"fee"`
	FeePayer        string                 `json:"feePayer"`
	Instructions    []HeliusInstruction    `json:"instructions"`
	NativeTransfers []HeliusNativeTransfer `json:"nativeTransfers"`
	Signature       string                 `json:"signature"`
	Slot            uint64                 `json:"slot"`
	Source          string                 `json:"source"`
	Timestamp       int64                  `json:"timestamp"`
	TokenTransfers  []HeliusTokenTransfer  `json:"tokenTransfers"`
	Type            string                 `json:"type"`
}

// HeliusInstruction Helius交易指令
type HeliusInstruction struct {
	Accounts          []string            `json:"accounts"`
	Data              string              `json:"data"` // Base58编码的指令数据
	ProgramId         string              `json:"programId"`
	InnerInstructions []HeliusInstruction `json:"innerInstructions"`
}

// HeliusNativeTransfer SOL原生转账
type HeliusNativeTransfer struct {
	Amount          uint64 `json:"amount"`
	FromUserAccount string `json:"fromUserAccount"`
	ToUserAccount   string `json:"toUserAccount"`
}

// HeliusTokenTransfer SPL Token转账
type HeliusTokenTransfer struct {
	FromTokenAccount string  `json:"fromTokenAccount"`
	FromUserAccount  string  `json:"fromUserAccount"`
	Mint             string  `json:"mint"`
	ToTokenAccount   string  `json:"toTokenAccount"`
	ToUserAccount    string  `json:"toUserAccount"`
	TokenAmount      float64 `json:"tokenAmount"`
	TokenStandard    string  `json:"tokenStandard"`
}

// HeliusAccountData 账户数据变更
type HeliusAccountData struct {
	Account             string `json:"account"`
	NativeBalanceChange int64  `json:"nativeBalanceChange"`
	TokenBalanceChanges []struct {
		Mint           string `json:"mint"`
		RawTokenAmount struct {
			Decimals    uint8  `json:"decimals"`
			TokenAmount string `json:"tokenAmount"`
		} `json:"rawTokenAmount"`
		TokenAccount string `json:"tokenAccount"`
		UserAccount  string `json:"userAccount"`
	} `json:"tokenBalanceChanges"`
}

// ==================== Solana RPC 响应结构 ====================

// SolanaBalanceResult getBalance 响应
type SolanaBalanceResult struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value uint64 `json:"value"`
}

// SolanaTokenAccountResult getTokenAccountsByOwner 响应
type SolanaTokenAccountResult struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value []struct {
		Account struct {
			Data struct {
				Parsed struct {
					Info struct {
						Mint        string `json:"mint"`
						Owner       string `json:"owner"`
						TokenAmount struct {
							Amount         string  `json:"amount"`
							Decimals       uint8   `json:"decimals"`
							UiAmount       float64 `json:"uiAmount"`
							UiAmountString string  `json:"uiAmountString"`
						} `json:"tokenAmount"`
					} `json:"info"`
					Type string `json:"type"`
				} `json:"parsed"`
				Program string `json:"program"`
				Space   int    `json:"space"`
			} `json:"data"`
			Executable bool   `json:"executable"`
			Lamports   uint64 `json:"lamports"`
			Owner      string `json:"owner"`
			RentEpoch  uint64 `json:"rentEpoch"`
		} `json:"account"`
		Pubkey string `json:"pubkey"`
	} `json:"value"`
}

// SolanaSignatureStatus getSignatureStatuses 响应
type SolanaSignatureStatus struct {
	Slot               uint64      `json:"slot"`
	Confirmations      *uint64     `json:"confirmations"`
	Err                interface{} `json:"err"`
	ConfirmationStatus string      `json:"confirmationStatus"`
}

// SolanaSignatureInfo getSignaturesForAddress 响应
type SolanaSignatureInfo struct {
	Signature          string      `json:"signature"`
	Slot               uint64      `json:"slot"`
	Err                interface{} `json:"err"`
	Memo               *string     `json:"memo"`
	BlockTime          *int64      `json:"blockTime"`
	ConfirmationStatus string      `json:"confirmationStatus"`
}

// SolanaBlockhashResult getLatestBlockhash 响应
type SolanaBlockhashResult struct {
	Context struct {
		Slot uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Blockhash            string `json:"blockhash"`
		LastValidBlockHeight uint64 `json:"lastValidBlockHeight"`
	} `json:"value"`
}
