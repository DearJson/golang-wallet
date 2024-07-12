package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
)

type BscBlock struct {
	Difficulty       string            `json:"difficulty"`
	ExtraData        string            `json:"extraData"`
	GasLimit         string            `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	Hash             string            `json:"hash"`
	LogsBloom        string            `json:"logsBloom"`
	Miner            string            `json:"miner"`
	MixHash          string            `json:"mixHash"`
	Nonce            string            `json:"nonce"`
	Number           string            `json:"number"`
	ParentHash       string            `json:"parentHash"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Sha3Uncles       string            `json:"sha3Uncles"`
	Size             string            `json:"size"`
	StateRoot        string            `json:"stateRoot"`
	Timestamp        string            `json:"timestamp"`
	TotalDifficulty  string            `json:"totalDifficulty"`
	Transactions     []BscTransactions `json:"transactions"`
	TransactionsRoot string            `json:"transactionsRoot"`
	Uncles           []interface{}     `json:"uncles"`
}
type BscTransactions struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}
type EthTransactions struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type BscTransferDetail struct {
	BlockHash        string      `json:"blockHash"`
	BlockNumber      string      `json:"blockNumber"`
	ContractAddress  interface{} `json:"contractAddress"`
	string           `json:"cumulativeGasUsed"`
	From             string                 `json:"from"`
	GasUsed          string                 `json:"gasUsed"`
	Logs             []BscTransferDetailLog `json:"logs"`
	LogsBloom        string                 `json:"logsBloom"`
	Status           string                 `json:"status"`
	To               string                 `json:"to"`
	TransactionHash  string                 `json:"transactionHash"`
	TransactionIndex string                 `json:"transactionIndex"`
	Type             string                 `json:"type"`
}
type EthTransferDetail struct {
	BlockHash         string                 `json:"blockHash"`
	BlockNumber       string                 `json:"blockNumber"`
	ContractAddress   interface{}            `json:"contractAddress"`
	CumulativeGasUsed string                 `json:"cumulativeGasUsed"`
	From              string                 `json:"from"`
	GasUsed           string                 `json:"gasUsed"`
	Logs              []BscTransferDetailLog `json:"logs"`
	LogsBloom         string                 `json:"logsBloom"`
	Status            string                 `json:"status"`
	To                string                 `json:"to"`
	TransactionHash   string                 `json:"transactionHash"`
	TransactionIndex  string                 `json:"transactionIndex"`
	Type              string                 `json:"type"`
}
type EthTransferDetailLog struct {
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
type BscTransferDetailLog struct {
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

type RechargeFunctionInputs struct {
	TokenAddress  []*common.Address  `json:"_tokenAddress"`
	Amount        []*decimal.Decimal `json:"_amount"`
	Remarks       string             `json:"_remarks"`
	CustomeCoin   []*common.Address  `json:"_customeCoin"`
	CustomeUser   []*common.Address  `json:"_customeUser"`
	CustomeAmount []*decimal.Decimal `json:"_customeAmount"`
}

// CardSynthesisCardInputs 卡牌合成struct
type CardSynthesisCardInputs struct {
	CardAddress   []*common.Address  `json:"_cardAddress"`
	CardTokenId   []*big.Int         `json:"_cardTokenId"`
	CardToAddress common.Address     `json:"_cardToAddress"`
	TokenAddress  []*common.Address  `json:"_tokenAddress"`
	Amount        []*decimal.Decimal `json:"_amount"`
	Remarks       string             `json:"_remarks"`
}

// dappHorc质押方法
type DappHorcDepositInputs struct {
	Amount         *big.Int       `json:"amount"`
	I              *big.Int       `json:"i"`
	Invitor        common.Address `json:"invitor"`
	MaxTokenAmount *big.Int       `json:"maxTokenAmount"`
}
