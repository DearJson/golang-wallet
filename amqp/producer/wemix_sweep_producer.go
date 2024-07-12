package producer

type WemixProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *WemixProducer) MsgContent() string {
	return b.Msg
}
