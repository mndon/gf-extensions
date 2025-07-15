// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLoginLog is the golang structure of table admin_login_log for DAO operations like Where/Data.
type AdminLoginLog struct {
	g.Meta        `orm:"table:admin_login_log, do:true"`
	InfoId        interface{} // 访问ID
	LoginName     interface{} // 登录账号
	Ipaddr        interface{} // 登录IP地址
	LoginLocation interface{} // 登录地点
	Browser       interface{} // 浏览器类型
	Os            interface{} // 操作系统
	Status        interface{} // 登录状态（0成功 1失败）
	Msg           interface{} // 提示消息
	LoginTime     *gtime.Time // 登录时间
	Module        interface{} // 登录模块
}
