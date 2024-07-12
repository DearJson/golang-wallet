// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/model/currency.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// Currency is the golang structure for table currency.
type Currency struct {
	Id              int64       `orm:"id,primary" json:"id"`                    //
	MainChain       string      `orm:"main_chain" json:"mainChain"`             // 主链
	Name            string      `orm:"name" json:"name"`                        // 币种名称
	ContractAddress string      `orm:"contract_address" json:"contractAddress"` // 合约地址
	Decimals        int         `orm:"decimals" json:"decimals"`                // 精度
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`             //
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`             //
}
