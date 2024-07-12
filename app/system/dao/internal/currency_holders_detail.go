// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/dao/internal/currency_holders_detail.go
// 生成人：gfast
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CurrencyHoldersDetailDao is the manager for logic model data accessing and custom defined data operations functions management.
type CurrencyHoldersDetailDao struct {
	Table   string                       // Table is the underlying table name of the DAO.
	Group   string                       // Group is the database configuration group name of current DAO.
	Columns CurrencyHoldersDetailColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CurrencyHoldersDetailColumns defines and stores column names for table currency_holders_detail.
type CurrencyHoldersDetailColumns struct {
	Id        string //
	HoldersId string // 跟踪记录ID
	Address   string // 地址
	Amount    string // 持币数量
	Per       string // 持币比例
	CreatedAt string //
	UpdatedAt string //
}

var currencyHoldersDetailColumns = CurrencyHoldersDetailColumns{
	Id:        "id",
	HoldersId: "holders_id",
	Address:   "address",
	Amount:    "amount",
	Per:       "per",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCurrencyHoldersDetailDao creates and returns a new DAO object for table data access.
func NewCurrencyHoldersDetailDao() *CurrencyHoldersDetailDao {
	return &CurrencyHoldersDetailDao{
		Group:   "default",
		Table:   "currency_holders_detail",
		Columns: currencyHoldersDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CurrencyHoldersDetailDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CurrencyHoldersDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CurrencyHoldersDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
