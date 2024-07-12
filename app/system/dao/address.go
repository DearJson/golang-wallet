// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 03:32:45
// 生成路径: gfast/app/system/dao/address.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// addressDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type addressDao struct {
	*internal.AddressDao
}

var (
	// Address is globally public accessible object for table tools_gen_table operations.
	Address = addressDao{
		internal.NewAddressDao(),
	}
)

// Fill with you ideas below.

// AddressSearchReq 分页请求参数
type AddressSearchReq struct {
	UserId    string `p:"userId"`    //用户标识
	MainChain string `p:"mainChain"` //主链
	Address   string `p:"address"`   //地址
	comModel.PageReq
}

// AddressAddReq 添加操作请求参数
type AddressAddReq struct {
	UserId     string `p:"userId" v:"required#用户标识不能为空"`
	MainChain  string `p:"mainChain"`
	Address    string `p:"address"`
	PrivateKey string `p:"privateKey"`
}

// AddressEditReq 修改操作请求参数
type AddressEditReq struct {
	Id         int    `p:"id" v:"required#主键ID不能为空"`
	UserId     int    `p:"userId" v:"required#用户标识不能为空"`
	MainChain  string `p:"mainChain" v:"required#主链不能为空"`
	Address    string `p:"address" v:"required#地址不能为空"`
	PrivateKey string `p:"privateKey" v:"required#私钥不能为空"`
}
