package producer

type EthProducer struct {
	Msg string
}

// MsgContent 实现以太坊扫块发送者
func (b *EthProducer) MsgContent() string {
	return b.Msg
}
