package producer

// SolanaProducer Solana扫块生产者
type SolanaProducer struct {
	Msg string
}

// MsgContent 实现Solana扫块发送者
func (t *SolanaProducer) MsgContent() string {
	return t.Msg
}
