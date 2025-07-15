// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDictTypeDao is the data access object for table admin_dict_type.
type AdminDictTypeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdminDictTypeColumns // columns contains all the column names of Table for convenient usage.
}

// AdminDictTypeColumns defines and stores column names for table admin_dict_type.
type AdminDictTypeColumns struct {
	DictId      string // 字典主键
	DictName    string // 字典名称
	DictType    string // 字典类型
	Status      string // 状态（0正常 1停用）
	CreateBy    string // 创建者
	UpdateBy    string // 更新者
	Remark      string // 备注
	CreatedTime string // 创建日期
	UpdatedTime string // 修改日期
}

// adminDictTypeColumns holds the columns for table admin_dict_type.
var adminDictTypeColumns = AdminDictTypeColumns{
	DictId:      "dict_id",
	DictName:    "dict_name",
	DictType:    "dict_type",
	Status:      "status",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
	Remark:      "remark",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}

// NewAdminDictTypeDao creates and returns a new DAO object for table data access.
func NewAdminDictTypeDao() *AdminDictTypeDao {
	return &AdminDictTypeDao{
		group:   "default",
		table:   "admin_dict_type",
		columns: adminDictTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminDictTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDictTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminDictTypeDao) Columns() AdminDictTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDictTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
