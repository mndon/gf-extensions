package eventbus

import (
	"context"
	"github.com/asaskevich/EventBus"
	"github.com/gogf/gf/v2/frame/g"
)

var Bus EventBus.Bus

func init() {
	Bus = EventBus.New()
	g.Log().Info(context.TODO(), "[eventbus] init")
}

// HasCallback
// @Description: 是否有订阅
// @param topic
// @return bool
func HasCallback(ctx context.Context, topic string) bool {
	return Bus.HasCallback(topic)
}

// WaitAsync
// @Description: 等待异步订阅执行全部执行完成
func WaitAsync(ctx context.Context) {
	Bus.WaitAsync()
}

// Subscribe
// @Description: 永久订阅
// @param topic
// @param fn
// @return error
func Subscribe(ctx context.Context, topic string, fn interface{}) error {
	g.Log().Infof(ctx, "[eventbus] subscribe topic: %s", topic)
	return Bus.Subscribe(topic, fn)
}

// SubscribeAsync
// @Description: 永久订阅，异步
// @param topic
// @param fn
// @param transactional
// @return error
func SubscribeAsync(ctx context.Context, topic string, fn interface{}, transactional bool) error {
	g.Log().Infof(ctx, "[eventbus] subscribe async topic: %s", topic)
	return Bus.SubscribeAsync(topic, fn, transactional)
}

// SubscribeOnce
// @Description: 订阅一次
// @param topic
// @param fn
// @return error
func SubscribeOnce(ctx context.Context, topic string, fn interface{}) error {
	g.Log().Infof(ctx, "[eventbus] subscribe once topic: %s", topic)
	return Bus.SubscribeOnce(topic, fn)
}

// SubscribeOnceAsync
// @Description: 订阅一次，异步
// @param topic
// @param fn
// @return error
func SubscribeOnceAsync(ctx context.Context, topic string, fn interface{}) error {
	g.Log().Infof(ctx, "[eventbus] subscribe once async topic: %s", topic)
	return Bus.SubscribeOnceAsync(topic, fn)
}

// Unsubscribe
// @Description: 取消订阅
// @param topic
// @param handler
// @return error
func Unsubscribe(ctx context.Context, topic string, handler interface{}) error {
	g.Log().Infof(ctx, "[eventbus] unsubscribe topic: %s", topic)
	return Bus.Unsubscribe(topic, handler)
}

// Publish
// @Description: 发布
// @param topic
// @param args
func Publish(ctx context.Context, topic string, args ...interface{}) {
	g.Log().Infof(ctx, "[eventbus] public topic: %s", topic)
	Bus.Publish(topic, args...)
}
