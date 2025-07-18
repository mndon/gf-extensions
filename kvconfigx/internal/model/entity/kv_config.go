// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KvConfig is the golang structure for table kv_config.
type KvConfig struct {
	Id            int64       `json:"id"             orm:"id"             ` //
	K             string      `json:"k"              orm:"k"              ` // key
	V             string      `json:"v"              orm:"v"              ` // value
	Description   string      `json:"description"    orm:"description"    ` // 描述
	ClientVisible int         `json:"client_visible" orm:"client_visible" ` // 客户端是否可见
	UpdatedTime   *gtime.Time `json:"updated_time"   orm:"updated_time"   ` // 更新时间
	CreatedTime   *gtime.Time `json:"created_time"   orm:"created_time"   ` // 添加时间
}
