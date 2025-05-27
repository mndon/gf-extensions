package delayqueuex

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/delayqueuex/internal/dao"
	"github.com/mndon/gf-extensions/delayqueuex/internal/do"
	"github.com/mndon/gf-extensions/delayqueuex/internal/entity"
	"github.com/mndon/gf-extensions/logx"
	"github.com/mndon/gf-extensions/slicex"
	"time"
)

type Message = entity.DelayQueue
type HandlerType func(ctx context.Context, r *Message) error

type Server struct {
	PollLimit    int
	PollInterval time.Duration
	handlerMap   map[string]HandlerType
}

func NewServer(pollLimit int, pollInterval time.Duration) *Server {
	return &Server{
		PollLimit:    pollLimit,
		PollInterval: pollInterval,
		handlerMap:   make(map[string]HandlerType),
	}
}

// AddTask 添加延迟任务
func (s *Server) AddTask(ctx context.Context, topic, body string, delaySeconds int) (int64, error) {
	readyTime := gtime.Now().Add(time.Duration(delaySeconds) * time.Second)

	id, err := dao.DelayQueue.Ctx(ctx).Data(do.DelayQueue{
		Topic:         topic,
		Body:          body,
		DelayDuration: delaySeconds,
		ReadyTime:     readyTime,
		Status:        StatusWait,
		RetryCount:    0,
	}).InsertAndGetId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Poll 获取就绪的任务
func (s *Server) Poll(ctx context.Context, topic string) (tasks []*entity.DelayQueue, err error) {
	// todo 分批拉取
	err = dao.DelayQueue.Ctx(ctx).Where(do.DelayQueue{Topic: topic, Status: StatusWait}).WhereLTE(dao.DelayQueue.Columns().ReadyTime, gtime.Now()).Limit(s.PollLimit).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return tasks, nil
	}

	// 更新任务状态为处理中
	ids := slicex.Map(tasks, func(item *entity.DelayQueue) int64 {
		return item.Id
	})
	_, err = dao.DelayQueue.Ctx(ctx).Data(do.DelayQueue{Status: StatusRunning}).WhereIn(dao.DelayQueue.Columns().Id, ids).Update()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// CompleteTask 标记任务完成
func (s *Server) CompleteTask(ctx context.Context, taskID int64, status int) error {
	_, err := dao.DelayQueue.Ctx(ctx).Data(do.DelayQueue{Status: status}).Where(dao.DelayQueue.Columns().Id, taskID).Update()
	if err != nil {
		return err
	}
	return nil
}

// RetryTask 重试任务
func (s *Server) RetryTask(ctx context.Context, taskID int64, delaySeconds int) error {
	readyTime := gtime.Now().Add(time.Duration(delaySeconds) * time.Second)

	_, err := dao.DelayQueue.Ctx(ctx).Where(do.DelayQueue{Id: taskID}).Data(
		do.DelayQueue{
			Status: StatusWait,
			RetryCount: &gdb.Counter{
				Field: dao.DelayQueue.Columns().RetryCount,
				Value: 1,
			},
			ReadyTime: readyTime,
		}).Update()
	if err != nil {
		return err
	}

	return nil
}

// Worker 工作协程
func (s *Server) Worker(ctx context.Context, topic string, handler HandlerType) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			tasks, err := s.Poll(ctx, topic)
			if err != nil {
				logx.New(ctx).Type("delay_queue_pool_err").Errorf("Poll tasks error: %+v", err)
				time.Sleep(s.PollInterval)
				continue
			}

			if len(tasks) == 0 {
				time.Sleep(s.PollInterval)
				continue
			}

			for _, task := range tasks {
				err = handler(gctx.New(), task)
				if err != nil {
					logx.New(ctx).Type("delay_queue_handle_err").Errorf("Handle task %d error: %+v", task.Id, err)
					// 重试逻辑
					if task.RetryCount < 2 {
						_ = s.RetryTask(ctx, task.Id, 10) // 10秒后重试
					} else {
						// 超过重试次数，标记为失败
						_ = s.CompleteTask(ctx, task.Id, StatusFail)
					}
					continue
				}

				// 处理成功
				err = s.CompleteTask(ctx, task.Id, StatusSuccess)
				if err != nil {
					logx.New(ctx).Type("delay_queue_complete_err").Errorf("Complete task %d error: %+v", task.Id, err)
				}
			}
		}
	}
}

// Register
// @Description: 注册队列消费handler
// @receiver s
// @param topic
// @param handler
func (s *Server) Register(topic string, handler HandlerType) {
	s.handlerMap[topic] = handler
}

// Run
// @Description: 运行
// @receiver s
// @param ctx
func (s *Server) Run() {
	for topic, handler := range s.handlerMap {
		go s.Worker(gctx.New(), topic, handler)
	}
}
