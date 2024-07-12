// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/dao/currency_holders_detail.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// currencyHoldersDetailDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type currencyHoldersDetailDao struct {
	*internal.CurrencyHoldersDetailDao
}

var (
	// CurrencyHoldersDetail is globally public accessible object for table tools_gen_table operations.
	CurrencyHoldersDetail = currencyHoldersDetailDao{
		internal.NewCurrencyHoldersDetailDao(),
	}
)

// Fill with you ideas below.

// CurrencyHoldersDetailSearchReq 分页请求参数
type CurrencyHoldersDetailSearchReq struct {
	HoldersId string  `p:"holdersId"` //跟踪记录ID
	Address   string  `p:"address"`   //地址
	Amount    float64 `p:"amount"`    //持币数量
	Per       float64 `p:"per"`       //持币比例
	comModel.PageReq
}

// CurrencyHoldersDetailAddReq 添加操作请求参数
type CurrencyHoldersDetailAddReq struct {
	HoldersId int     `p:"holdersId" v:"required#跟踪记录ID不能为空"`
	Address   string  `p:"address" v:"required#地址不能为空"`
	Amount    float64 `p:"amount" v:"required#持币数量不能为空"`
	Per       float64 `p:"per" v:"required#持币比例不能为空"`
}

// CurrencyHoldersDetailEditReq 修改操作请求参数
type CurrencyHoldersDetailEditReq struct {
	Id        int     `p:"id" v:"required#主键ID不能为空"`
	HoldersId int     `p:"holdersId" v:"required#跟踪记录ID不能为空"`
	Address   string  `p:"address" v:"required#地址不能为空"`
	Amount    float64 `p:"amount" v:"required#持币数量不能为空"`
	Per       float64 `p:"per" v:"required#持币比例不能为空"`
}
