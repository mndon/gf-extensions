// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure of table admin_role for DAO operations like Where/Data.
type AdminRole struct {
	g.Meta      `orm:"table:admin_role, do:true"`
	Id          interface{} //
	Status      interface{} // 状态;0:禁用;1:正常
	ListOrder   interface{} // 排序
	Name        interface{} // 角色名称
	Remark      interface{} // 备注
	DataScope   interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedTime *gtime.Time // 创建日期
	UpdatedTime *gtime.Time // 修改日期
}
