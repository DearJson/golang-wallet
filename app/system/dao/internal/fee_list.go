// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/internal/fee_list.go
// 生成人：gfast
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// FeeListDao is the manager for logic model data accessing and custom defined data operations functions management.
type FeeListDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns FeeListColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// FeeListColumns defines and stores column names for table fee_list.
type FeeListColumns struct {    
    Id  string  //    
    MainChain  string  // 主链    
    CoinName  string  // 手续费币种    
    Address  string  // 地址    
    Amount  string  // 转移手续费    
    Hash  string  // hash    
    Status  string  // 状态 1-待上链 2-上链成功 3-上链失败    
    RechargeId  string  // 充值ID    
    CreatedAt  string  //    
    UpdatedAt  string  //    
}
var feeListColumns = FeeListColumns{    
    Id:  "id",    
    MainChain:  "main_chain",    
    CoinName:  "coin_name",    
    Address:  "address",    
    Amount:  "amount",    
    Hash:  "hash",    
    Status:  "status",    
    RechargeId:  "recharge_id",    
    CreatedAt:  "created_at",    
    UpdatedAt:  "updated_at",    
}
// NewFeeListDao creates and returns a new DAO object for table data access.
func NewFeeListDao() *FeeListDao {
	return &FeeListDao{
        Group:    "default",
        Table: "fee_list",
        Columns:feeListColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeeListDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeeListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeeListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}