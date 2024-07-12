package producer

import "math/big"

type TronProducer struct {
	Msg string
}

type TronTransfer struct {
	BlockHash   string   `json:"block_hash"`
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Amount      *big.Int `json:"amount"`
}

// MsgContent 实现波场扫块发送者
func (t *TronProducer) MsgContent() string {
	return t.Msg
}
