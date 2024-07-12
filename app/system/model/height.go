// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/model/height.go
// 生成人：gfast
// ==========================================================================

package model

// Height is the golang structure for table height.
type Height struct {
	Id        uint64 `orm:"id,primary" json:"id"`        //
	MainChain string `orm:"main_chain" json:"mainChain"` // 主网
	Height    int    `orm:"height" json:"height"`        // 块号
}
