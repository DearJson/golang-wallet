package rpc

type TronTransaction struct {
	Transaction []byte `json:"transaction"`
	BlockHeader int64  `json:"blockHeader"`
	BlockHash   string `json:"blockHash"`
}
