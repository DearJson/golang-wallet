// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/internal/currency.go
// 生成人：gfast
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// CurrencyDao is the manager for logic model data accessing and custom defined data operations functions management.
type CurrencyDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns CurrencyColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// CurrencyColumns defines and stores column names for table currency.
type CurrencyColumns struct {    
    Id  string  //    
    MainChain  string  // 主链    
    Name  string  // 币种名称    
    ContractAddress  string  // 合约地址    
    Decimals  string  // 精度    
    CreatedAt  string  //    
    UpdatedAt  string  //    
}
var currencyColumns = CurrencyColumns{    
    Id:  "id",    
    MainChain:  "main_chain",    
    Name:  "name",    
    ContractAddress:  "contract_address",    
    Decimals:  "decimals",    
    CreatedAt:  "created_at",    
    UpdatedAt:  "updated_at",    
}
// NewCurrencyDao creates and returns a new DAO object for table data access.
func NewCurrencyDao() *CurrencyDao {
	return &CurrencyDao{
        Group:    "default",
        Table: "currency",
        Columns:currencyColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CurrencyDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CurrencyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CurrencyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}