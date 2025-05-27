package delayqueuex

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/delayqueuex/internal/dao"
	"github.com/mndon/gf-extensions/delayqueuex/internal/do"
	"time"
)

var defaultServer = NewServer(100, time.Second*10)

// Run
// @Description: 初始化
// @param ctx
func Run() {
	defaultServer.Run()
}

// AddTask
// @Description: 添加任务
// @param ctx
// @param topic
// @param body
// @param delaySeconds
// @return int64
// @return error
func AddTask(ctx context.Context, topic, body string, delaySeconds int) (int64, error) {
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

// Register
// @Description: 注册队列消费handler
// @receiver s
// @param topic
// @param handler
func Register(topic string, handler HandlerType) {
	defaultServer.handlerMap[topic] = handler
}
