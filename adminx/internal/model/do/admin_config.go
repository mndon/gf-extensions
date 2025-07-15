// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminConfig is the golang structure of table admin_config for DAO operations like Where/Data.
type AdminConfig struct {
	g.Meta      `orm:"table:admin_config, do:true"`
	ConfigId    interface{} // 参数主键
	ConfigName  interface{} // 参数名称
	ConfigKey   interface{} // 参数键名
	ConfigValue interface{} // 参数键值
	ConfigType  interface{} // 系统内置（Y是 N否）
	CreateBy    interface{} // 创建者
	UpdateBy    interface{} // 更新者
	Remark      interface{} // 备注
	CreatedTime *gtime.Time // 创建日期
	UpdatedTime *gtime.Time // 修改日期
}
