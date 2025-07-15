// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDictType is the golang structure of table admin_dict_type for DAO operations like Where/Data.
type AdminDictType struct {
	g.Meta      `orm:"table:admin_dict_type, do:true"`
	DictId      interface{} // 字典主键
	DictName    interface{} // 字典名称
	DictType    interface{} // 字典类型
	Status      interface{} // 状态（0正常 1停用）
	CreateBy    interface{} // 创建者
	UpdateBy    interface{} // 更新者
	Remark      interface{} // 备注
	CreatedTime *gtime.Time // 创建日期
	UpdatedTime *gtime.Time // 修改日期
}
