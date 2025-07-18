// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// KvConfig is the golang structure of table kv_config for DAO operations like Where/Data.
type KvConfig struct {
	g.Meta        `orm:"table:kv_config, do:true"`
	Id            interface{} //
	K             interface{} // key
	V             interface{} // value
	Description   interface{} // 描述
	ClientVisible interface{} // 客户端是否可见
	UpdatedTime   *gtime.Time // 更新时间
	CreatedTime   *gtime.Time // 添加时间
}
