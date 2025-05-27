// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DelayQueue is the golang structure of table delay_queue for DAO operations like Where/Data.
type DelayQueue struct {
	g.Meta        `orm:"table:delay_queue, do:true"`
	Id            interface{} //
	Topic         interface{} // 队列主题/分类
	Body          interface{} // 任务内容，通常为JSON
	DelayDuration interface{} // 延迟时长(秒)
	ReadyTime     *gtime.Time // 就绪时间(创建时间 + 延迟时间)
	Status        interface{} // 状态: 0-等待中, 1-处理中, 2-已完成, 3-已取消
	RetryCount    interface{} // 重试次数
	UpdatedTime   *gtime.Time // 更新时间
	CreatedTime   *gtime.Time // 添加时间
}
