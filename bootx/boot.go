// Package bootx 提供额外的启动逻辑，如：pprof性能分析
package bootx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/logx"
)

func ExtBootUp(ctx context.Context, s *ghttp.Server) {
	// 1. pprof性能分析, 参考：https://goframe.org/docs/web/senior-pprof, 默认会自动注册以下几个路由规则：
	// /debug/pprof/*action
	// /debug/pprof/cmdline
	// /debug/pprof/profile
	// /debug/pprof/symbol
	// /debug/pprof/trace
	if g.Cfg().MustGet(ctx, "pprof.enable", false).Bool() {
		s.EnablePProf()
	}

	// 2. 日志滚动
	if g.Cfg().MustGet(ctx, "logger.logxRotateEnable").Bool() {
		logx.NewRotate(
			g.Cfg().MustGet(ctx, "logger.path").String(),
			g.Cfg().MustGet(ctx, "logger.logxRotateCountLimit").Int(),
			g.Cfg().MustGet(ctx, "logger.logxRotateCheckInterval").String(),
		).RotateChecksTimely(ctx)
	}

	if g.Cfg().MustGet(ctx, "database.logger.logxRotateEnable").Bool() {
		logx.NewRotate(
			g.Cfg().MustGet(ctx, "database.logger.path").String(),
			g.Cfg().MustGet(ctx, "database.logger.logxRotateCountLimit").Int(),
			g.Cfg().MustGet(ctx, "database.logger.logxRotateCheckInterval").String(),
		).RotateChecksTimely(ctx)
	}
}
