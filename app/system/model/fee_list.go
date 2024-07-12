// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/model/fee_list.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// FeeList is the golang structure for table fee_list.
type FeeList struct {
	Id         uint64      `orm:"id,primary" json:"id"`          //
	MainChain  string      `orm:"main_chain" json:"mainChain"`   // 主链
	CoinName   string      `orm:"coin_name" json:"coinName"`     // 手续费币种
	Address    string      `orm:"address" json:"address"`        // 地址
	Amount     float64     `orm:"amount" json:"amount"`          // 转移手续费
	Hash       string      `orm:"hash" json:"hash"`              // hash
	Status     int8        `orm:"status" json:"status"`          // 状态 1-待上链 2-上链成功 3-上链失败
	RechargeId uint64      `orm:"recharge_id" json:"rechargeId"` // 充值ID
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"`   //
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"`   //
}
