// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/dao/currency_holders.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// currencyHoldersDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type currencyHoldersDao struct {
	*internal.CurrencyHoldersDao
}

var (
	// CurrencyHolders is globally public accessible object for table tools_gen_table operations.
	CurrencyHolders = currencyHoldersDao{
		internal.NewCurrencyHoldersDao(),
	}
)

// Fill with you ideas below.

// CurrencyHoldersSearchReq 分页请求参数
type CurrencyHoldersSearchReq struct {
	MainChain       string `p:"mainChain"`       //链
	ContractAddress string `p:"contractAddress"` //合约地址
	Decimals        int8   `p:"decimals"`        //精度
	comModel.PageReq
}

// CurrencyHoldersAddReq 添加操作请求参数
type CurrencyHoldersAddReq struct {
	MainChain       string `p:"mainChain" v:"required#链不能为空"`
	ContractAddress string `p:"contractAddress" v:"required#合约地址不能为空"`
	Decimals        int8   `p:"decimals" v:"required#精度不能为空"`
}

// CurrencyHoldersEditReq 修改操作请求参数
type CurrencyHoldersEditReq struct {
	Id              int    `p:"id" v:"required#主键ID不能为空"`
	MainChain       string `p:"mainChain" v:"required#链不能为空"`
	ContractAddress string `p:"contractAddress" v:"required#合约地址不能为空"`
	Decimals        int8   `p:"decimals" v:"required#精度不能为空"`
}
