package producer

type NacProducer struct {
	Msg string
}

// MsgContent 实现NAC扫块发送者
func (b *NacProducer) MsgContent() string {
	return b.Msg
}
