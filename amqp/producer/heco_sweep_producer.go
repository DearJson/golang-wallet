package producer

type HecoProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *HecoProducer) MsgContent() string {
	return b.Msg
}
