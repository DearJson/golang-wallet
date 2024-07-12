package model

import "time"

type MonitorSwapTrade struct {
	ID               int64     `orm:"column:id" db:"id"`                               //  id主键
	MainChain        string    `orm:"column:main_chain" db:"main_chain"`               //  主链
	BlockHash        string    `orm:"column:block_hash" db:"block_hash"`               //  块hash
	FromAddress      string    `orm:"column:from_address" db:"from_address"`           //  交易方地址
	ContractAddress  string    `orm:"column:contract_address" db:"contract_address"`   //  兑换前币种
	ContractAddress1 string    `orm:"column:contract_address1" db:"contract_address1"` //  兑换后币种
	Amount           string    `orm:"column:amount" db:"amount"`                       //  兑换前数量
	Amount1          string    `orm:"column:amount1" db:"amount1"`                     //  兑换后数量
	Character        string    `orm:"column:character" db:"character"`                 //  兑换路径
	ToAddress        string    `orm:"column:to_address" db:"to_address"`               //  收币地址
	Type             int64     `orm:"column:type" db:"type"`                           //  类型 1-买入 2-卖出
	Hash             string    `orm:"column:hash" db:"hash"`                           //  交易hash
	BlockHeight      string    `orm:"column:block_height" db:"block_height"`           //  区块高度
	CallNumber       int64     `orm:"column:call_number" db:"call_number"`             //  回调次数
	CreatedAt        time.Time `orm:"column:created_at" db:"created_at"`
	UpdatedAt        time.Time `orm:"column:updated_at" db:"updated_at"`
}
