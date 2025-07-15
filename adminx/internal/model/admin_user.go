/*
* @desc:用户模型对象
* @company:
* @Author:
* @Date:   2022/3/7 11:47
 */

package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
)

// LoginUserRes 登录返回
type LoginUserRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`           //
	UserName     string `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string `orm:"user_password"    json:"userPassword"` // 登录密码;cmf_password加密
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`      // 是否后台管理员 1 是  0   否
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
	DeptId       uint64 `orm:"dept_id"       json:"deptId"`          //部门id
}

// AdminUserRoleDeptRes 带有部门、角色、岗位信息的用户数据
type AdminUserRoleDeptRes struct {
	*entity.AdminUser
	RoleInfo []*AdminUserRoleInfoRes `json:"roleInfo"`
	Post     []*AdminUserPostInfoRes `json:"post"`
}

type AdminUserRoleInfoRes struct {
	RoleId uint   `json:"roleId"`
	Name   string `json:"name"`
}

type AdminUserPostInfoRes struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
}

type AdminUserSimpleRes struct {
	gmeta.Meta   `orm:"table:admin_user"`
	Id           uint64 `orm:"id"       json:"id"`                   //
	Avatar       string `orm:"avatar" json:"avatar"`                 // 头像
	Sex          int    `orm:"sex" json:"sex"`                       // 性别
	UserName     string `orm:"user_name" json:"userName"`            // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
}
