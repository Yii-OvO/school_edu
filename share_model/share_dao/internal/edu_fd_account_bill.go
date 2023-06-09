// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EduFdAccountBillDao is the data access object for table edu_fd_account_bill.
type EduFdAccountBillDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns EduFdAccountBillColumns // columns contains all the column names of Table for convenient usage.
}

// EduFdAccountBillColumns defines and stores column names for table edu_fd_account_bill.
type EduFdAccountBillColumns struct {
	Id            string // ID
	FromUserId    string // 交易发起方UserID，如果是系统则固定为-1
	ToUserId      string // 交易对象UserID
	FdAccountId   string // 财务账户ID
	BeforeBalance string // 交易前账户余额
	Amount        string // 交易金额
	AfterBalance  string // 交易后账户余额
	UnionOrderId  string // 关联业务订单ID
	InOutType     string // 收支类型：1收入，2支出
	TradeType     string // 交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它
	TradeAt       string // 交易时间
	Remark        string // 备注信息
	TradeState    string // 交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
	DeletedAt     string //
	CreatedAt     string //
	CreatedBy     string //
}

// eduFdAccountBillColumns holds the columns for table edu_fd_account_bill.
var eduFdAccountBillColumns = EduFdAccountBillColumns{
	Id:            "id",
	FromUserId:    "from_user_id",
	ToUserId:      "to_user_id",
	FdAccountId:   "fd_account_id",
	BeforeBalance: "before_balance",
	Amount:        "amount",
	AfterBalance:  "after_balance",
	UnionOrderId:  "union_order_id",
	InOutType:     "in_out_type",
	TradeType:     "trade_type",
	TradeAt:       "trade_at",
	Remark:        "remark",
	TradeState:    "trade_state",
	DeletedAt:     "deleted_at",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
}

// NewEduFdAccountBillDao creates and returns a new DAO object for table data access.
func NewEduFdAccountBillDao(proxy ...dao_interface.IDao) *EduFdAccountBillDao {
	var dao *EduFdAccountBillDao
	if len(proxy) > 0 {
		dao = &EduFdAccountBillDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: eduFdAccountBillColumns,
		}
		return dao
	}

	return &EduFdAccountBillDao{
		group:   "default",
		table:   "edu_fd_account_bill",
		columns: eduFdAccountBillColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EduFdAccountBillDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EduFdAccountBillDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *EduFdAccountBillDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *EduFdAccountBillDao) Columns() EduFdAccountBillColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EduFdAccountBillDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *EduFdAccountBillDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EduFdAccountBillDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}