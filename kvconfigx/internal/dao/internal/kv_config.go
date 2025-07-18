// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KvConfigDao is the data access object for table kv_config.
type KvConfigDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns KvConfigColumns // columns contains all the column names of Table for convenient usage.
}

// KvConfigColumns defines and stores column names for table kv_config.
type KvConfigColumns struct {
	Id            string //
	K             string // key
	V             string // value
	Description   string // 描述
	ClientVisible string // 客户端是否可见
	UpdatedTime   string // 更新时间
	CreatedTime   string // 添加时间
}

// kvConfigColumns holds the columns for table kv_config.
var kvConfigColumns = KvConfigColumns{
	Id:            "id",
	K:             "k",
	V:             "v",
	Description:   "description",
	ClientVisible: "client_visible",
	UpdatedTime:   "updated_time",
	CreatedTime:   "created_time",
}

// NewKvConfigDao creates and returns a new DAO object for table data access.
func NewKvConfigDao() *KvConfigDao {
	return &KvConfigDao{
		group:   "default",
		table:   "kv_config",
		columns: kvConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KvConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KvConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KvConfigDao) Columns() KvConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KvConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KvConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KvConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
