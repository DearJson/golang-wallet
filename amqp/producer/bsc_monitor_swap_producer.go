package producer

type BscMonitorSwapProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *BscMonitorSwapProducer) MsgContent() string {
	return b.Msg
}
