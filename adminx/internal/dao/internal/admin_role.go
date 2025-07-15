// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleDao is the data access object for table admin_role.
type AdminRoleDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminRoleColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleColumns defines and stores column names for table admin_role.
type AdminRoleColumns struct {
	Id          string //
	Status      string // 状态;0:禁用;1:正常
	ListOrder   string // 排序
	Name        string // 角色名称
	Remark      string // 备注
	DataScope   string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedTime string // 创建日期
	UpdatedTime string // 修改日期
}

// adminRoleColumns holds the columns for table admin_role.
var adminRoleColumns = AdminRoleColumns{
	Id:          "id",
	Status:      "status",
	ListOrder:   "list_order",
	Name:        "name",
	Remark:      "remark",
	DataScope:   "data_scope",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// NewAdminRoleDao creates and returns a new DAO object for table data access.
func NewAdminRoleDao() *AdminRoleDao {
	return &AdminRoleDao{
		group:   "default",
		table:   "admin_role",
		columns: adminRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleDao) Columns() AdminRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
