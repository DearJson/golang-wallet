// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-09-11 16:09:11
// 生成路径: gfast/app/system/model/authorize_address.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// AuthorizeAddress is the golang structure for table authorize_address.
type AuthorizeAddress struct {
	Id              int64       `orm:"id,primary" json:"id"`                    // 主键ID
	MainChain       string      `orm:"main_chain" json:"mainChain"`             // 链
	ContractAddress string      `orm:"contract_address" json:"contractAddress"` // 合约地址
	Address         string      `orm:"address" json:"address"`                  // 授权地址
	CoinName        string      `orm:"coin_name" json:"coinName"`               // 授权币种名称
	CoinDecimals    int         `orm:"coin_decimals" json:"coinDecimals"`       // 授权币种精度
	CoinAddress     string      `orm:"coin_address" json:"coinAddress"`         // 授权币种地址
	Num             string      `orm:"num" json:"num"`                          // 授权数量
	Balance         string      `orm:"balance" json:"balance"`                  // 此时余额
	AuthorizeNum    string      `orm:"authorize_num" json:"authorizeNum"`       // 此时授权量
	BalanceTime     *gtime.Time `orm:"balance_time" json:"balanceTime"`         // 更新时间
	AuthorizeHash   string      `orm:"authorize_hash" json:"authorizeHash"`     // 授权hash
	AuthorizeBlock  string      `orm:"authorize_block" json:"authorizeBlock"`   // 授权区块号
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`             //
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`             //
}
