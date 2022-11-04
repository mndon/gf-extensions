package http

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SafeFiltering 安全限制异常返回
func SafeFiltering(ctx context.Context) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeSecurityReason.Code(),
		Message: gcode.CodeSecurityReason.Message(),
	})
	r.ExitAll()
}
