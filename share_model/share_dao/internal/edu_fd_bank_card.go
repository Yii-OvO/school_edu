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

// EduFdBankCardDao is the data access object for table edu_fd_bank_card.
type EduFdBankCardDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns EduFdBankCardColumns // columns contains all the column names of Table for convenient usage.
}

// EduFdBankCardColumns defines and stores column names for table edu_fd_bank_card.
type EduFdBankCardColumns struct {
	Id            string // ID
	BankName      string // 银行名称
	CardType      string // 银行卡类型：1借记卡，2储蓄卡
	CardNumber    string // 银行卡号
	ExpiredAt     string // 有效期
	HolderName    string // 银行卡开户名
	BankOfAccount string // 开户行
	State         string // 状态：0禁用，1正常
	Remark        string // 备注信息
	UserId        string // 用户id，表示属于谁
	CreatedAt     string //
	CreatedBy     string //
	UpdatedAt     string //
	UpdatedBy     string //
	DeletedAt     string //
	DeletedBy     string //
}

// eduFdBankCardColumns holds the columns for table edu_fd_bank_card.
var eduFdBankCardColumns = EduFdBankCardColumns{
	Id:            "id",
	BankName:      "bank_name",
	CardType:      "card_type",
	CardNumber:    "card_number",
	ExpiredAt:     "expired_at",
	HolderName:    "holder_name",
	BankOfAccount: "bank_of_account",
	State:         "state",
	Remark:        "remark",
	UserId:        "user_id",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
	DeletedAt:     "deleted_at",
	DeletedBy:     "deleted_by",
}

// NewEduFdBankCardDao creates and returns a new DAO object for table data access.
func NewEduFdBankCardDao(proxy ...dao_interface.IDao) *EduFdBankCardDao {
	var dao *EduFdBankCardDao
	if len(proxy) > 0 {
		dao = &EduFdBankCardDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: eduFdBankCardColumns,
		}
		return dao
	}

	return &EduFdBankCardDao{
		group:   "default",
		table:   "edu_fd_bank_card",
		columns: eduFdBankCardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EduFdBankCardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EduFdBankCardDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *EduFdBankCardDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *EduFdBankCardDao) Columns() EduFdBankCardColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EduFdBankCardDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *EduFdBankCardDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
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
func (dao *EduFdBankCardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
