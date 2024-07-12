// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 03:32:45
// 生成路径: gfast/app/system/model/address.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// Address is the golang structure for table address.
type Address struct {
	Id         uint64      `orm:"id,primary" json:"id"`          //
	UserId     string      `orm:"user_id" json:"userId"`         // 用户标识
	MainChain  string      `orm:"main_chain" json:"mainChain"`   // 主链
	Address    string      `orm:"address" json:"address"`        // 地址
	PrivateKey string      `orm:"private_key" json:"privateKey"` // 私钥
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"`   //
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"`   //
}
