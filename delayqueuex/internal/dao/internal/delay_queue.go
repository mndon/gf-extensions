// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DelayQueueDao is the data access object for table delay_queue.
type DelayQueueDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DelayQueueColumns // columns contains all the column names of Table for convenient usage.
}

// DelayQueueColumns defines and stores column names for table delay_queue.
type DelayQueueColumns struct {
	Id            string //
	Topic         string // 队列主题/分类
	Body          string // 任务内容，通常为JSON
	DelayDuration string // 延迟时长(秒)
	ReadyTime     string // 就绪时间(创建时间 + 延迟时间)
	Status        string // 状态: 0-等待中, 1-处理中, 2-已完成, 3-已取消
	RetryCount    string // 重试次数
	UpdatedTime   string // 更新时间
	CreatedTime   string // 添加时间
}

// delayQueueColumns holds the columns for table delay_queue.
var delayQueueColumns = DelayQueueColumns{
	Id:            "id",
	Topic:         "topic",
	Body:          "body",
	DelayDuration: "delay_duration",
	ReadyTime:     "ready_time",
	Status:        "status",
	RetryCount:    "retry_count",
	UpdatedTime:   "updated_time",
	CreatedTime:   "created_time",
}

// NewDelayQueueDao creates and returns a new DAO object for table data access.
func NewDelayQueueDao() *DelayQueueDao {
	return &DelayQueueDao{
		group:   "default",
		table:   "delay_queue",
		columns: delayQueueColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DelayQueueDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DelayQueueDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DelayQueueDao) Columns() DelayQueueColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DelayQueueDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DelayQueueDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DelayQueueDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
