// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserDao is the data access object for table admin_user.
type AdminUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AdminUserColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUserColumns defines and stores column names for table admin_user.
type AdminUserColumns struct {
	Id            string //
	UserName      string // 用户名
	Mobile        string // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string // 用户昵称
	Birthday      string // 生日
	UserPassword  string // 登录密码;cmf_password加密
	UserSalt      string // 加密盐
	UserStatus    string // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string // 用户登录邮箱
	Sex           string // 性别;0:保密,1:男,2:女
	Avatar        string // 用户头像
	DeptId        string // 部门id
	Remark        string // 备注
	IsAdmin       string // 是否后台管理员 1 是  0   否
	Address       string // 联系地址
	Describe      string // 描述信息
	LastLoginIp   string // 最后登录ip
	LastLoginTime string // 最后登录时间
	CreatedTime   string // 创建日期
	UpdatedTime   string // 修改日期
	DeletedAt     string // 删除时间
}

// adminUserColumns holds the columns for table admin_user.
var adminUserColumns = AdminUserColumns{
	Id:            "id",
	UserName:      "user_name",
	Mobile:        "mobile",
	UserNickname:  "user_nickname",
	Birthday:      "birthday",
	UserPassword:  "user_password",
	UserSalt:      "user_salt",
	UserStatus:    "user_status",
	UserEmail:     "user_email",
	Sex:           "sex",
	Avatar:        "avatar",
	DeptId:        "dept_id",
	Remark:        "remark",
	IsAdmin:       "is_admin",
	Address:       "address",
	Describe:      "describe",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedTime:   "created_time",
	UpdatedTime:   "updated_time",
	DeletedAt:     "deleted_at",
}

// NewAdminUserDao creates and returns a new DAO object for table data access.
func NewAdminUserDao() *AdminUserDao {
	return &AdminUserDao{
		group:   "default",
		table:   "admin_user",
		columns: adminUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminUserDao) Columns() AdminUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
