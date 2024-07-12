// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:34
// 生成路径: gfast/app/system/dao/internal/withdraw.go
// 生成人：gfast
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// WithdrawDao is the manager for logic model data accessing and custom defined data operations functions management.
type WithdrawDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns WithdrawColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// WithdrawColumns defines and stores column names for table withdraw.
type WithdrawColumns struct {
	Id                 string //
	MainChain          string // 主链
	CoinToken          string // 币种
	Address            string // 转出地址
	Amount             string // 提币数量
	ContractAddress    string // 合约地址
	Status             string // 状态 1待审核，2待上链，3上链中，4上链失败，5上链成功，6已拒绝
	Nonce              string // nonce
	ConfirmationNumber string // 上链确认次数
	Hash               string // 交易Hash
	Remarks            string // 交易备注
	TrxRemark          string //trx交易备注
	NotifyUrl          string // 交易回调地址
	CreatedAt          string //
	UpdatedAt          string //
}

var withdrawColumns = WithdrawColumns{
	Id:                 "id",
	MainChain:          "main_chain",
	CoinToken:          "coin_token",
	Address:            "address",
	Amount:             "amount",
	ContractAddress:    "contract_address",
	Status:             "status",
	Nonce:              "nonce",
	ConfirmationNumber: "confirmation_number",
	Hash:               "hash",
	Remarks:            "remarks",
	TrxRemark:          "trx_remark",
	NotifyUrl:          "notify_url",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewWithdrawDao creates and returns a new DAO object for table data access.
func NewWithdrawDao() *WithdrawDao {
	return &WithdrawDao{
		Group:   "default",
		Table:   "withdraw",
		Columns: withdrawColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WithdrawDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WithdrawDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
