package producer

type BscAuthorizeProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *BscAuthorizeProducer) MsgContent() string {
	return b.Msg
}
