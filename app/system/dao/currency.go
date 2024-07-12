// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/currency.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// currencyDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type currencyDao struct {
	*internal.CurrencyDao
}

var (
	// Currency is globally public accessible object for table tools_gen_table operations.
	Currency = currencyDao{
		internal.NewCurrencyDao(),
	}
)

// Fill with you ideas below.

// CurrencySearchReq 分页请求参数
type CurrencySearchReq struct {
	MainChain       string `p:"mainChain"`       //主链
	Name            string `p:"name"`            //币种名称
	ContractAddress string `p:"contractAddress"` //合约地址
	Decimals        string `p:"decimals"`        //精度
	comModel.PageReq
}

// CurrencyAddReq 添加操作请求参数
type CurrencyAddReq struct {
	MainChain       string `p:"mainChain"`
	Name            string `p:"name" v:"required#币种名称不能为空"`
	ContractAddress string `p:"contractAddress" `
	Decimals        int    `p:"decimals" v:"required#精度不能为空"`
}

// CurrencyEditReq 修改操作请求参数
type CurrencyEditReq struct {
	Id              int64  `p:"id" v:"required#主键ID不能为空"`
	MainChain       string `p:"mainChain" v:"required#主链不能为空"`
	Name            string `p:"name" v:"required#币种名称不能为空"`
	ContractAddress string `p:"contractAddress" `
	Decimals        int    `p:"decimals" v:"required#精度不能为空"`
}
