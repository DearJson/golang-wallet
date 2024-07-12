// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/internal/height.go
// 生成人：gfast
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// HeightDao is the manager for logic model data accessing and custom defined data operations functions management.
type HeightDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns HeightColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// HeightColumns defines and stores column names for table height.
type HeightColumns struct {    
    Id  string  //    
    MainChain  string  // 主网    
    Height  string  // 块号    
}
var heightColumns = HeightColumns{    
    Id:  "id",    
    MainChain:  "main_chain",    
    Height:  "height",    
}
// NewHeightDao creates and returns a new DAO object for table data access.
func NewHeightDao() *HeightDao {
	return &HeightDao{
        Group:    "default",
        Table: "height",
        Columns:heightColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HeightDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HeightDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HeightDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}