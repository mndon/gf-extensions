// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUserOnline is the golang structure of table admin_user_online for DAO operations like Where/Data.
type AdminUserOnline struct {
	g.Meta     `orm:"table:admin_user_online, do:true"`
	Id         interface{} //
	Uuid       interface{} // 用户标识
	Token      interface{} // 用户token
	CreateTime *gtime.Time // 登录时间
	UserName   interface{} // 用户名
	Ip         interface{} // 登录ip
	Explorer   interface{} // 浏览器
	Os         interface{} // 操作系统
}
