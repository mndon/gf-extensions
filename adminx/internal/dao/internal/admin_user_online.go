// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserOnlineDao is the data access object for table admin_user_online.
type AdminUserOnlineDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns AdminUserOnlineColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUserOnlineColumns defines and stores column names for table admin_user_online.
type AdminUserOnlineColumns struct {
	Id         string //
	Uuid       string // 用户标识
	Token      string // 用户token
	CreateTime string // 登录时间
	UserName   string // 用户名
	Ip         string // 登录ip
	Explorer   string // 浏览器
	Os         string // 操作系统
}

// adminUserOnlineColumns holds the columns for table admin_user_online.
var adminUserOnlineColumns = AdminUserOnlineColumns{
	Id:         "id",
	Uuid:       "uuid",
	Token:      "token",
	CreateTime: "create_time",
	UserName:   "user_name",
	Ip:         "ip",
	Explorer:   "explorer",
	Os:         "os",
}

// NewAdminUserOnlineDao creates and returns a new DAO object for table data access.
func NewAdminUserOnlineDao() *AdminUserOnlineDao {
	return &AdminUserOnlineDao{
		group:   "default",
		table:   "admin_user_online",
		columns: adminUserOnlineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminUserOnlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminUserOnlineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminUserOnlineDao) Columns() AdminUserOnlineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminUserOnlineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminUserOnlineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminUserOnlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
