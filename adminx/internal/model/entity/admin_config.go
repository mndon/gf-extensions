// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminConfig is the golang structure for table admin_config.
type AdminConfig struct {
	ConfigId    uint        `json:"configId"    orm:"config_id"    ` // 参数主键
	ConfigName  string      `json:"configName"  orm:"config_name"  ` // 参数名称
	ConfigKey   string      `json:"configKey"   orm:"config_key"   ` // 参数键名
	ConfigValue string      `json:"configValue" orm:"config_value" ` // 参数键值
	ConfigType  int         `json:"configType"  orm:"config_type"  ` // 系统内置（Y是 N否）
	CreateBy    uint        `json:"createBy"    orm:"create_by"    ` // 创建者
	UpdateBy    uint        `json:"updateBy"    orm:"update_by"    ` // 更新者
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
	CreatedTime *gtime.Time `json:"createdTime" orm:"created_time" ` // 创建日期
	UpdatedTime *gtime.Time `json:"updatedTime" orm:"updated_time" ` // 修改日期
}
