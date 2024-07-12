// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/model/currency_holders.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// CurrencyHolders is the golang structure for table currency_holders.
type CurrencyHolders struct {
	Id              int         `orm:"id,primary" json:"id"`                    //
	MainChain       string      `orm:"main_chain" json:"mainChain"`             // 链
	ContractAddress string      `orm:"contract_address" json:"contractAddress"` // 合约地址
	Decimals        int8        `orm:"decimals" json:"decimals"`                // 精度
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`             //
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`             //
}
