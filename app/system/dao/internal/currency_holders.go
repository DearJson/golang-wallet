// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/dao/internal/currency_holders.go
// 生成人：gfast
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CurrencyHoldersDao is the manager for logic model data accessing and custom defined data operations functions management.
type CurrencyHoldersDao struct {
	Table   string                 // Table is the underlying table name of the DAO.
	Group   string                 // Group is the database configuration group name of current DAO.
	Columns CurrencyHoldersColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CurrencyHoldersColumns defines and stores column names for table currency_holders.
type CurrencyHoldersColumns struct {
	Id              string //
	MainChain       string // 链
	ContractAddress string // 合约地址
	Decimals        string // 精度
	CreatedAt       string //
	UpdatedAt       string //
}

var currencyHoldersColumns = CurrencyHoldersColumns{
	Id:              "id",
	MainChain:       "main_chain",
	ContractAddress: "contract_address",
	Decimals:        "decimals",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewCurrencyHoldersDao creates and returns a new DAO object for table data access.
func NewCurrencyHoldersDao() *CurrencyHoldersDao {
	return &CurrencyHoldersDao{
		Group:   "default",
		Table:   "currency_holders",
		Columns: currencyHoldersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CurrencyHoldersDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CurrencyHoldersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CurrencyHoldersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
