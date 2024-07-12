// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/recharge.go
// 生成人：gfast
// ==========================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// rechargeDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type rechargeDao struct {
	*internal.RechargeDao
}

var (
	// Recharge is globally public accessible object for table tools_gen_table operations.
	Recharge = rechargeDao{
		internal.NewRechargeDao(),
	}
)

// Fill with you ideas below.

// RechargeSearchReq 分页请求参数
type RechargeSearchReq struct {
	MainChain        string `p:"mainChain"`        //主链
	BlockHash        string `p:"blockHash"`        //块hash
	FromAddress      string `p:"fromAddress"`      //发送方地址
	ToAddress        string `p:"toAddress"`        //充币地址
	ContractAddress  string `p:"contractAddress"`  //币种1合约地址
	ContractAddress1 string `p:"contractAddress1"` //币种2合约地址
	Hash             string `p:"hash"`             //充币交易hashId
	BlockHeight      string `p:"blockHeight"`      //区块高度
	Status           string `p:"status"`           //状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中
	ImputationHash   string `p:"imputationHash"`   //归集hash
	RechargeType     int8   `p:"rechargeType"`     //充币类型
	TimeRange        string `p:"timeRange"`        //时间筛选
	comModel.PageReq
}

// RechargeAddReq 添加操作请求参数
type RechargeAddReq struct {
	MainChain        string `p:"mainChain" `
	BlockHash        string `p:"blockHash" `
	CoinToken        string `p:"coinToken" `
	CoinToken1       string `p:"coinToken1" `
	FromAddress      string `p:"fromAddress" `
	ToAddress        string `p:"toAddress" `
	Amount           string `p:"amount" `
	Amount1          string `p:"amount1" `
	ContractAddress  string `p:"contractAddress" `
	ContractAddress1 string `p:"contractAddress1" `
	Hash             string `p:"hash" v:"required#充币交易hashId不能为空"`
	BlockHeight      string `p:"blockHeight" `
	Status           int8   `p:"status" v:"required#状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中不能为空"`
	Remarks          string `p:"remarks" `
	RechargeType     int8   `p:"rechargeType" v:"required#充值类型 不能为空"`
	TokenId          string `p:"tokenId"`
	CustomeUser      string `p:"customeUser"`
	CustomeCoin      string `p:"customeCoin"`
	CustomeAmount    string `p:"customeAmount"`
}

// RechargeEditReq 修改操作请求参数
type RechargeEditReq struct {
	Id               int64   `p:"id" v:"required#主键ID不能为空"`
	MainChain        string  `p:"mainChain" `
	BlockHash        string  `p:"blockHash" `
	CoinToken        string  `p:"coinToken" `
	CoinToken1       string  `p:"coinToken1" `
	FromAddress      string  `p:"fromAddress" `
	ToAddress        string  `p:"toAddress" `
	Amount           float64 `p:"amount" `
	Amount1          float64 `p:"amount1" `
	ContractAddress  string  `p:"contractAddress" `
	ContractAddress1 string  `p:"contractAddress1" `
	Hash             string  `p:"hash" v:"required#充币交易hashId不能为空"`
	BlockHeight      string  `p:"blockHeight" `
	CallNumber       int     `p:"callNumber" `
	Status           int     `p:"status" v:"required#状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中不能为空"`
	ImputationHash   string  `p:"imputationHash" `
	Remarks          string  `p:"remarks" `
	RechargeType     int8    `p:"rechargeType" v:"required#充值类型 不能为空"`
}

// RechargeStatusReq 设置用户状态参数
type RechargeStatusReq struct {
	Id     int64 `p:"id" v:"required#主键ID不能为空"`
	Status int   `p:"status" v:"required#状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中不能为空"`
}
