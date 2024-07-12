package producer

type BscProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *BscProducer) MsgContent() string {
	return b.Msg
}
