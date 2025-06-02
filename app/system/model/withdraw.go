// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:34
// 生成路径: gfast/app/system/model/withdraw.go
// 生成人：gfast
// ==========================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Withdraw is the golang structure for table withdraw.
type Withdraw struct {
	Id                 uint64      `orm:"id,primary" json:"id"`                          //
	MainChain          string      `orm:"main_chain" json:"mainChain"`                   // 主链
	CoinToken          string      `orm:"coin_token" json:"coinToken"`                   // 币种
	Address            string      `orm:"address" json:"address"`                        // 转出地址
	Amount             float64     `orm:"amount" json:"amount"`                          // 提币数量
	ContractAddress    string      `orm:"contract_address" json:"contractAddress"`       // 合约地址
	Status             int8        `orm:"status" json:"status"`                          // 状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝
	Nonce              uint64      `orm:"nonce" json:"nonce"`                            // nonce
	ConfirmationNumber int         `orm:"confirmation_number" json:"confirmationNumber"` // 上链确认次数
	Hash               string      `orm:"hash" json:"hash"`                              // 交易Hash
	Remarks            string      `orm:"remarks" json:"remarks"`                        // 交易备注
	NotifyUrl          string      `orm:"notify_url" json:"notifyUrl"`                   // 交易回调地址
	TokenId            int64       `orm:"token_id" json:"token_id"`                      //tokenID (nft使用)
	Url                string      `orm:"url" json:"url"`                                //url (nft使用)
	HashKey            string      `orm:"hashKey" json:"hashKey"`
	Nonce1             string      `orm:"nonce1" json:"nonce1"`
	Currency           *Currency   `orm:"with:contract_address=contract_address" json:"currency"` //关联币种
	SwapRoute          string      `orm:"swap_route" json:"swap_route"`
	SwapPath           string      `orm:"swap_path" json:"swap_path"`
	TrxRemark          string      `orm:"trx_remark" json:"trx_remark"`
	Function           string      `orm:"function" json:"function"`
	FunctionAddress    string      `orm:"function_address" json:"function_address"`
	CreatedAt          *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt          *gtime.Time `orm:"updated_at" json:"updatedAt"` //
}
