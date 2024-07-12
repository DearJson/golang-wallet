// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 03:32:45
// 生成路径: gfast/app/system/dao/internal/address.go
// 生成人：gfast
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// AddressDao is the manager for logic model data accessing and custom defined data operations functions management.
type AddressDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns AddressColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// AddressColumns defines and stores column names for table address.
type AddressColumns struct {    
    Id  string  //    
    UserId  string  // 用户标识    
    MainChain  string  // 主链    
    Address  string  // 地址    
    PrivateKey  string  // 私钥    
    CreatedAt  string  //    
    UpdatedAt  string  //    
}
var addressColumns = AddressColumns{    
    Id:  "id",    
    UserId:  "user_id",    
    MainChain:  "main_chain",    
    Address:  "address",    
    PrivateKey:  "private_key",    
    CreatedAt:  "created_at",    
    UpdatedAt:  "updated_at",    
}
// NewAddressDao creates and returns a new DAO object for table data access.
func NewAddressDao() *AddressDao {
	return &AddressDao{
        Group:    "default",
        Table: "address",
        Columns:addressColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddressDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}