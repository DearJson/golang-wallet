// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-09-11 16:09:11
// 生成路径: gfast/app/system/dao/authorize_address.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
	"github.com/gogf/gf/os/gtime"
)

// authorizeAddressDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type authorizeAddressDao struct {
	*internal.AuthorizeAddressDao
}

var (
	// AuthorizeAddress is globally public accessible object for table tools_gen_table operations.
	AuthorizeAddress = authorizeAddressDao{
		internal.NewAuthorizeAddressDao(),
	}
)

// Fill with you ideas below.

// AuthorizeAddressSearchReq 分页请求参数
type AuthorizeAddressSearchReq struct {
	MainChain       string      `p:"mainChain"`       //链
	ContractAddress string      `p:"contractAddress"` //合约地址
	Address         string      `p:"address"`         //授权地址
	CoinName        string      `p:"coinName"`        //授权币种名称
	CoinDecimals    string      `p:"coinDecimals"`    //授权币种精度
	CoinAddress     string      `p:"coinAddress"`     //授权币种地址
	Num             string      `p:"num"`             //授权数量
	Balance         string      `p:"balance"`         //此时余额
	AuthorizeNum    string      `p:"authorizeNum"`    //此时授权量
	BalanceTime     *gtime.Time `p:"balanceTime"`     //更新时间
	AuthorizeHash   string      `p:"authorizeHash"`   //授权hash
	AuthorizeBlock  string      `p:"authorizeBlock"`  //授权区块号
	comModel.PageReq
}

// AuthorizeAddressAddReq 添加操作请求参数
type AuthorizeAddressAddReq struct {
	MainChain       string `p:"mainChain" v:"required#链不能为空"`
	ContractAddress string `p:"contractAddress" v:"required#合约地址不能为空"`
	Address         string `p:"address" v:"required#授权地址不能为空"`
	CoinName        string `p:"coinName" v:"required#授权币种名称不能为空"`
	CoinDecimals    int    `p:"coinDecimals" `
	CoinAddress     string `p:"coinAddress" `
	Num             string `p:"num" `
	//Balance  string   `p:"balance" `
	//AuthorizeNum  string   `p:"authorizeNum" `
	//BalanceTime  *gtime.Time   `p:"balanceTime" `
	AuthorizeHash  string `p:"authorizeHash" `
	AuthorizeBlock string `p:"authorizeBlock" `
}

// AuthorizeAddressEditReq 修改操作请求参数
type AuthorizeAddressEditReq struct {
	Id              int64       `p:"id" v:"required#主键ID不能为空"`
	MainChain       string      `p:"mainChain" v:"required#链不能为空"`
	ContractAddress string      `p:"contractAddress" v:"required#合约地址不能为空"`
	Address         string      `p:"address" v:"required#授权地址不能为空"`
	CoinName        string      `p:"coinName" v:"required#授权币种名称不能为空"`
	CoinDecimals    int         `p:"coinDecimals" `
	CoinAddress     string      `p:"coinAddress" `
	Num             string      `p:"num" `
	Balance         string      `p:"balance" `
	AuthorizeNum    string      `p:"authorizeNum" `
	BalanceTime     *gtime.Time `p:"balanceTime" `
	AuthorizeHash   string      `p:"authorizeHash" `
	AuthorizeBlock  string      `p:"authorizeBlock" `
}
