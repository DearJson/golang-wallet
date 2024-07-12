// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/fee_list.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// feeListDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type feeListDao struct {
	*internal.FeeListDao
}

var (
	// FeeList is globally public accessible object for table tools_gen_table operations.
	FeeList = feeListDao{
		internal.NewFeeListDao(),
	}
)

// Fill with you ideas below.

// FeeListSearchReq 分页请求参数
type FeeListSearchReq struct {
	MainChain  string `p:"mainChain"`  //主链
	Address    string `p:"address"`    //地址
	Hash       string `p:"hash"`       //hash
	Status     int8   `p:"status"`     //状态 1-待上链 2-上链成功 3-上链失败
	RechargeId uint64 `p:"rechargeId"` //充值ID
	comModel.PageReq
}

// FeeListAddReq 添加操作请求参数
type FeeListAddReq struct {
	MainChain  string  `p:"mainChain" v:"required#主链不能为空"`
	CoinName   string  `p:"coinName" v:"required#手续费币种不能为空"`
	Address    string  `p:"address" v:"required#地址不能为空"`
	Amount     float64 `p:"amount" v:"required#转移手续费不能为空"`
	Hash       string  `p:"hash" v:"required#hash不能为空"`
	Status     int8    `p:"status" v:"required#状态 1-待上链 2-上链成功 3-上链失败不能为空"`
	RechargeId uint64  `p:"rechargeId" v:"required#充值ID不能为空"`
}

// FeeListEditReq 修改操作请求参数
type FeeListEditReq struct {
	Id         uint64  `p:"id" v:"required#主键ID不能为空"`
	MainChain  string  `p:"mainChain" v:"required#主链不能为空"`
	CoinName   string  `p:"coinName" v:"required#手续费币种不能为空"`
	Address    string  `p:"address" v:"required#地址不能为空"`
	Amount     float64 `p:"amount" v:"required#转移手续费不能为空"`
	Hash       string  `p:"hash" v:"required#hash不能为空"`
	Status     int8    `p:"status" v:"required#状态 1-待上链 2-上链成功 3-上链失败不能为空"`
	RechargeId uint64  `p:"rechargeId" v:"required#充值ID不能为空"`
}

// FeeListStatusReq 设置用户状态参数
type FeeListStatusReq struct {
	Id     uint64 `p:"id" v:"required#主键ID不能为空"`
	Status int8   `p:"status" v:"required#状态 1-待上链 2-上链成功 3-上链失败不能为空"`
}
