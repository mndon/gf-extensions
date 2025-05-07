package logx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	HandleTypePlain = "plain"
	HandleTypeJSON  = "json"
)

// Init
// @Description: logx初始化
//  1. 配置输出格式，配置项：
//     logger.logxHandle：handle名称，plain-可视化文本 json-json格式
//  2. 配置日志滚动删除，配置项：
//     logger.logxRotateEnable：是否开启日志滚动删除
//     logger.path：日志路径
//     logger.logxRotateCountLimit：保留几份日志
//     logger.logxRotateCheckInterval： 滚动脚本运行频率
func Init(ctx context.Context) {
	handleType, _ := g.Cfg().Get(ctx, "logger.logxHandle", "text")
	switch handleType.String() {
	case HandleTypePlain:
		g.Log().SetHandlers(HandlerPlain)
	case HandleTypeJSON:
		g.Log().SetHandlers(HandlerJson)
	default:
		g.Log().SetHandlers(HandlerPlain)
	}

	// 日志滚动
	if g.Cfg().MustGet(ctx, "logger.logxRotateEnable", false).Bool() {
		NewRotate(
			g.Cfg().MustGet(ctx, "logger.path").String(),
			g.Cfg().MustGet(ctx, "logger.logxRotateCountLimit").Int(),
			g.Cfg().MustGet(ctx, "logger.logxRotateCheckInterval").String(),
		).RotateChecksTimely(ctx)
	}

	if g.Cfg().MustGet(ctx, "database.logger.logxRotateEnable", false).Bool() {
		NewRotate(
			g.Cfg().MustGet(ctx, "database.logger.path").String(),
			g.Cfg().MustGet(ctx, "database.logger.logxRotateCountLimit").Int(),
			g.Cfg().MustGet(ctx, "database.logger.logxRotateCheckInterval").String(),
		).RotateChecksTimely(ctx)
	}
}
