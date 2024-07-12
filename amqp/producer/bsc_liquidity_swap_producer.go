package producer

type BscLiquiditySwapProducer struct {
	Msg string
}

// MsgContent 实现币安扫块发送者
func (b *BscLiquiditySwapProducer) MsgContent() string {
	return b.Msg
}
