// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/model/recharge.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// Recharge is the golang structure for table recharge.
type Recharge struct {
	Id               int64  `orm:"id,primary" json:"id"`                      // id主键
	MainChain        string `orm:"main_chain" json:"mainChain"`               // 主链
	BlockHash        string `orm:"block_hash" json:"blockHash"`               // 块hash
	CoinToken        string `orm:"coin_token" json:"coinToken"`               // 币种1
	CoinToken1       string `orm:"coin_token1" json:"coinToken1"`             // 币种2
	FromAddress      string `orm:"from_address" json:"fromAddress"`           // 发送方地址
	ToAddress        string `orm:"to_address" json:"toAddress"`               // 充币地址
	Amount           string `orm:"amount" json:"amount"`                      // 充币数量
	Amount1          string `orm:"amount1" json:"amount1"`                    // 充币2数量
	ContractAddress  string `orm:"contract_address" json:"contractAddress"`   // 币种1合约地址
	ContractAddress1 string `orm:"contract_address1" json:"contractAddress1"` // 币种2合约地址
	Hash             string `orm:"hash" json:"hash"`                          // 充币交易hashId
	BlockHeight      string `orm:"block_height" json:"blockHeight"`           // 区块高度
	CallNumber       int    `orm:"call_number" json:"callNumber"`             // 回调次数
	Status           int    `orm:"status" json:"status"`                      // 状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中
	ImputationHash   string `orm:"imputation_hash" json:"imputationHash"`     // 归集hash
	Remarks          string `orm:"remarks" json:"remarks"`                    // 备注
	RechargeType     int8   `orm:"recharge_type" json:"rechargeType"`         // 充值类型 1-单币充值 2-双币充值
	TokenId          string `orm:"token_id" json:"tokenId"`

	CustomeUser   string `orm:"customeUser"`
	CustomeCoin   string `orm:"customeCoin"`
	CustomeAmount string `orm:"customeAmount"`

	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
}
