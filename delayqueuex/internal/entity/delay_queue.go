// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DelayQueue is the golang structure for table delay_queue.
type DelayQueue struct {
	Id            int64       `json:"id"             orm:"id"             ` //
	Topic         string      `json:"topic"          orm:"topic"          ` // 队列主题/分类
	Body          string      `json:"body"           orm:"body"           ` // 任务内容，通常为JSON
	DelayDuration int         `json:"delay_duration" orm:"delay_duration" ` // 延迟时长(秒)
	ReadyTime     *gtime.Time `json:"ready_time"     orm:"ready_time"     ` // 就绪时间(创建时间 + 延迟时间)
	Status        int         `json:"status"         orm:"status"         ` // 状态: 0-等待中, 1-处理中, 2-已完成, 3-已取消
	RetryCount    int         `json:"retry_count"    orm:"retry_count"    ` // 重试次数
	UpdatedTime   *gtime.Time `json:"updated_time"   orm:"updated_time"   ` // 更新时间
	CreatedTime   *gtime.Time `json:"created_time"   orm:"created_time"   ` // 添加时间
}
