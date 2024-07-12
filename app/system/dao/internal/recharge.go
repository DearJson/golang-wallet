// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/internal/recharge.go
// 生成人：gfast
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// RechargeDao is the manager for logic model data accessing and custom defined data operations functions management.
type RechargeDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns RechargeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// RechargeColumns defines and stores column names for table recharge.
type RechargeColumns struct {    
    Id  string  // id主键    
    MainChain  string  // 主链    
    BlockHash  string  // 块hash    
    CoinToken  string  // 币种1    
    CoinToken1  string  // 币种2    
    FromAddress  string  // 发送方地址    
    ToAddress  string  // 充币地址    
    Amount  string  // 充币数量    
    Amount1  string  // 充币2数量    
    ContractAddress  string  // 币种1合约地址    
    ContractAddress1  string  // 币种2合约地址    
    Hash  string  // 充币交易hashId    
    BlockHeight  string  // 区块高度    
    CallNumber  string  // 回调次数    
    Status  string  // 状态，1充币成功待归集，2归集上链中，3归集成功, 4归集失败，5充值费用中    
    ImputationHash  string  // 归集hash    
    Remarks  string  // 备注    
    RechargeType  string  // 充币数量    
    CreatedAt  string  //    
    UpdatedAt  string  //    
}
var rechargeColumns = RechargeColumns{    
    Id:  "id",    
    MainChain:  "main_chain",    
    BlockHash:  "block_hash",    
    CoinToken:  "coin_token",    
    CoinToken1:  "coin_token1",    
    FromAddress:  "from_address",    
    ToAddress:  "to_address",    
    Amount:  "amount",    
    Amount1:  "amount1",    
    ContractAddress:  "contract_address",    
    ContractAddress1:  "contract_address1",    
    Hash:  "hash",    
    BlockHeight:  "block_height",    
    CallNumber:  "call_number",    
    Status:  "status",    
    ImputationHash:  "imputation_hash",    
    Remarks:  "remarks",    
    RechargeType:  "recharge_type",    
    CreatedAt:  "created_at",    
    UpdatedAt:  "updated_at",    
}
// NewRechargeDao creates and returns a new DAO object for table data access.
func NewRechargeDao() *RechargeDao {
	return &RechargeDao{
        Group:    "default",
        Table: "recharge",
        Columns:rechargeColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RechargeDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RechargeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}