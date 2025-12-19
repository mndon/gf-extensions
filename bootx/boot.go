// Package bootx 提供额外的启动逻辑，如：pprof性能分析
package bootx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/configx"
	"github.com/mndon/gf-extensions/logx"
	"github.com/mndon/gf-extensions/uidx"
)

func init() {
	// 1. 配置初始化
	configx.Init()

	// 2. 日志配置
	logx.Init(context.Background())

	// 3. 雪花发号器初始化
	uidx.Init(0)
}

// HttpServerBootUp
// @Description: 初始化http服务
// @param ctx
// @param s
func HttpServerBootUp(ctx context.Context, s *ghttp.Server) {
	// 1. pprof性能分析, 参考：https://goframe.org/docs/web/senior-pprof, 默认会自动注册以下几个路由规则：
	// /debug/pprof/*action
	// /debug/pprof/cmdline
	// /debug/pprof/profile
	// /debug/pprof/symbol
	// /debug/pprof/trace
	if g.Cfg().MustGet(ctx, "pprof.enable", false).Bool() {
		s.EnablePProf()
	}

	if g.Cfg().MustGet(ctx, "pprof.enable", false).Bool() {
		enableHealthCheck(s)
	}

}

// TestBootUp
// @Description: 初始化单元测试
// @param ctx
func TestBootUp(ctx context.Context) {
	// 当前未有需要初始化的逻辑， 但是还是需要引入，从而达到执行初始化函数init的目的
}
