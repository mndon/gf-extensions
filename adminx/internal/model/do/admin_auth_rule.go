// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminAuthRule is the golang structure of table admin_auth_rule for DAO operations like Where/Data.
type AdminAuthRule struct {
	g.Meta      `orm:"table:admin_auth_rule, do:true"`
	Id          interface{} //
	Pid         interface{} // 父ID
	Name        interface{} // 规则名称
	Title       interface{} // 规则名称
	Icon        interface{} // 图标
	Condition   interface{} // 条件
	Remark      interface{} // 备注
	MenuType    interface{} // 类型 0目录 1菜单 2按钮
	Weigh       interface{} // 权重
	IsHide      interface{} // 显示状态
	Path        interface{} // 路由地址
	Component   interface{} // 组件路径
	IsLink      interface{} // 是否外链 1是 0否
	ModuleType  interface{} // 所属模块
	ModelId     interface{} // 模型ID
	IsIframe    interface{} // 是否内嵌iframe
	IsCached    interface{} // 是否缓存
	Redirect    interface{} // 路由重定向地址
	IsAffix     interface{} // 是否固定
	LinkUrl     interface{} // 链接地址
	CreatedTime *gtime.Time // 创建日期
	UpdatedTime *gtime.Time // 修改日期
}
