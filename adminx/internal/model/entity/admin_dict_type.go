// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDictType is the golang structure for table admin_dict_type.
type AdminDictType struct {
	DictId      uint64      `json:"dictId"      orm:"dict_id"      ` // 字典主键
	DictName    string      `json:"dictName"    orm:"dict_name"    ` // 字典名称
	DictType    string      `json:"dictType"    orm:"dict_type"    ` // 字典类型
	Status      uint        `json:"status"      orm:"status"       ` // 状态（0正常 1停用）
	CreateBy    uint        `json:"createBy"    orm:"create_by"    ` // 创建者
	UpdateBy    uint        `json:"updateBy"    orm:"update_by"    ` // 更新者
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` // 创建日期
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" ` // 修改日期
}
