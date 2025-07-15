// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	Id          uint        `json:"id"          orm:"id"           ` //
	Status      uint        `json:"status"      orm:"status"       ` // 状态;0:禁用;1:正常
	ListOrder   uint        `json:"listOrder"   orm:"list_order"   ` // 排序
	Name        string      `json:"name"        orm:"name"         ` // 角色名称
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
	DataScope   uint        `json:"dataScope"   orm:"data_scope"   ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` // 创建日期
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" ` // 修改日期
}
