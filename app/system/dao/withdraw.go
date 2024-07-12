// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 04:18:34
// 生成路径: gfast/app/system/dao/withdraw.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// withdrawDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type withdrawDao struct {
	*internal.WithdrawDao
}

var (
	// Withdraw is globally public accessible object for table tools_gen_table operations.
	Withdraw = withdrawDao{
		internal.NewWithdrawDao(),
	}
)

// Fill with you ideas below.

// WithdrawSearchReq 分页请求参数
type WithdrawSearchReq struct {
	MainChain          string  `p:"mainChain"`          //主链
	CoinToken          string  `p:"coinToken"`          //币种
	Address            string  `p:"address"`            //转出地址
	Amount             float64 `p:"amount"`             //提币数量
	ContractAddress    string  `p:"contractAddress"`    //合约地址
	Status             int8    `p:"status"`             //状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝
	Nonce              uint64  `p:"nonce"`              //nonce
	ConfirmationNumber int     `p:"confirmationNumber"` //上链确认次数
	Hash               string  `p:"hash"`               //交易Hash
	Remarks            string  `p:"remarks"`            //交易备注
	NotifyUrl          string  `p:"notifyUrl"`          //交易回调地址
	TimeRange          string  `p:"timeRange"`          //时间筛选
	comModel.PageReq
}

// WithdrawAddReq 添加操作请求参数
type WithdrawAddReq struct {
	MainChain       string  `p:"mainChain"`
	CoinToken       string  `p:"coinToken"`
	Sign            string  `p:"sign"`
	Address         string  `p:"address" v:"required#转出地址不能为空"`
	Amount          float64 `p:"amount" v:"required#提币数量不能为空"`
	ContractAddress string  `p:"contract_address" v:"required#提现币种合约地址不能为空"`
	NotifyUrl       string  `p:"notify_url" `
	Remarks         string  `p:"remarks" `
	Status          int8    `p:"status"`
	HashKey         string  `P:"hashKey"`
	Nonce1          string  `p:"nonce1"`
	SwapRoute       string  `p:"swap_route"`
	SwapPath        string  `p:"swap_path"`
	TrxRemark       string  `p:"trx_remark"`
}

// WithdrawNftAddReq 添加NFT操作请求参数
type WithdrawNftAddReq struct {
	MainChain       string  `p:"mainChain"`
	CoinToken       string  `p:"coinToken"`
	Address         string  `p:"address" v:"required#转出地址不能为空"`
	Amount          float64 `p:"amount" v:"required#提现数量不能为空"`
	ContractAddress string  `p:"contract_address" v:"required#提现币种合约地址不能为空"`
	TokenId         string  `p:"token_id" v:"required#token_id不能为空"`
	Url             string  `p:"url" `
	NotifyUrl       string  `p:"notify_url" `
	Remarks         string  `p:"remarks" `
	Status          int8    `p:"status"`
	HashKey         string  `P:"hashKey"`
	Nonce1          string  `p:"nonce1"`
}

// WithdrawEditReq 修改操作请求参数
type WithdrawEditReq struct {
	Id                 uint64  `p:"id" v:"required#主键ID不能为空"`
	MainChain          string  `p:"mainChain" v:"required#主链不能为空"`
	CoinToken          string  `p:"coinToken" v:"required#币种 不能为空"`
	Address            string  `p:"address" v:"required#转出地址不能为空"`
	Amount             float64 `p:"amount" v:"required#提币数量不能为空"`
	ContractAddress    string  `p:"contractAddress" `
	Status             int8    `p:"status" v:"required#状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝不能为空"`
	Nonce              uint64  `p:"nonce" v:"required#nonce不能为空"`
	ConfirmationNumber int     `p:"confirmationNumber" v:"required#上链确认次数不能为空"`
	Hash               string  `p:"hash" `
	Remarks            string  `p:"remarks" `
	NotifyUrl          string  `p:"notifyUrl" `
}

// WithdrawStatusReq 设置用户状态参数
type WithdrawStatusReq struct {
	Id     uint64 `p:"id" v:"required#主键ID不能为空"`
	Status int8   `p:"status" v:"required#状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝不能为空"`
}

// LpInfoReq LP合约地址
type LpInfoReq struct {
	ContractAddress string `p:"contract_address" v:"required#合约地址不能为空"`
}

type BalanceOfReq struct {
	ContractAddress string `p:"contract_address" v:"required#合约地址不能为空"`
	Address         string `p:"address" v:"required#地址不能为空"`
}

type SetAddressReq struct {
	UserId  string `p:"user_id" v:"required#用户标识不能为空"`
	Address string `p:"address" v:"required#地址不能为空"`
}
