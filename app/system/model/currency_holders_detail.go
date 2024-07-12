// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/model/currency_holders_detail.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// CurrencyHoldersDetail is the golang structure for table currency_holders_detail.
type CurrencyHoldersDetail struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	HoldersId int         `orm:"holders_id" json:"holdersId"` // 跟踪记录ID
	Address   string      `orm:"address" json:"address"`      // 地址
	Amount    float64     `orm:"amount" json:"amount"`        // 持币数量
	Per       float64     `orm:"per" json:"per"`              // 持币比例
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
}
