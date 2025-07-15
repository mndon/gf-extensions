// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminLoginLogDao is the data access object for table admin_login_log.
type AdminLoginLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdminLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// AdminLoginLogColumns defines and stores column names for table admin_login_log.
type AdminLoginLogColumns struct {
	InfoId        string // 访问ID
	LoginName     string // 登录账号
	Ipaddr        string // 登录IP地址
	LoginLocation string // 登录地点
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Status        string // 登录状态（0成功 1失败）
	Msg           string // 提示消息
	LoginTime     string // 登录时间
	Module        string // 登录模块
}

// adminLoginLogColumns holds the columns for table admin_login_log.
var adminLoginLogColumns = AdminLoginLogColumns{
	InfoId:        "info_id",
	LoginName:     "login_name",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginTime:     "login_time",
	Module:        "module",
}

// NewAdminLoginLogDao creates and returns a new DAO object for table data access.
func NewAdminLoginLogDao() *AdminLoginLogDao {
	return &AdminLoginLogDao{
		group:   "default",
		table:   "admin_login_log",
		columns: adminLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminLoginLogDao) Columns() AdminLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
