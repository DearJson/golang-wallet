// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-09-11 16:09:11
// 生成路径: gfast/app/system/dao/internal/authorize_address.go
// 生成人：gfast
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// AuthorizeAddressDao is the manager for logic model data accessing and custom defined data operations functions management.
type AuthorizeAddressDao struct {
	Table   string                  // Table is the underlying table name of the DAO.
	Group   string                  // Group is the database configuration group name of current DAO.
	Columns AuthorizeAddressColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AuthorizeAddressColumns defines and stores column names for table authorize_address.
type AuthorizeAddressColumns struct {
	Id              string // 主键ID
	MainChain       string // 链
	ContractAddress string // 合约地址
	Address         string // 授权地址
	CoinName        string // 授权币种名称
	CoinDecimals    string // 授权币种精度
	CoinAddress     string // 授权币种地址
	Num             string // 授权数量
	Balance         string // 此时余额
	AuthorizeNum    string // 此时授权量
	BalanceTime     string // 更新时间
	AuthorizeHash   string // 授权hash
	AuthorizeBlock  string // 授权区块号
	CreatedAt       string //
	UpdatedAt       string //
}

var authorizeAddressColumns = AuthorizeAddressColumns{
	Id:              "id",
	MainChain:       "main_chain",
	ContractAddress: "contract_address",
	Address:         "address",
	CoinName:        "coin_name",
	CoinDecimals:    "coin_decimals",
	CoinAddress:     "coin_address",
	Num:             "num",
	Balance:         "balance",
	AuthorizeNum:    "authorize_num",
	BalanceTime:     "balance_time",
	AuthorizeHash:   "authorize_hash",
	AuthorizeBlock:  "authorize_block",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewAuthorizeAddressDao creates and returns a new DAO object for table data access.
func NewAuthorizeAddressDao() *AuthorizeAddressDao {
	return &AuthorizeAddressDao{
		Group:   "default",
		Table:   "authorize_address",
		Columns: authorizeAddressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthorizeAddressDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthorizeAddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthorizeAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
