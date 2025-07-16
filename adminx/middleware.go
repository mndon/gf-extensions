package adminx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

// MiddlewareLoginAuth
// @Description: 登陆鉴权中间件
// @param r
func MiddlewareLoginAuth(r *ghttp.Request) {
	b, _ := service.GfToken().IsLogin(r)
	if !b {
		r.SetError(gerror.NewCode(gcode.New(4010, "token已失效", nil)))
		return
	}
	r.Middleware.Next()
}

// MiddlewareCtx
// @Description: 用户信息上下文
// @param r
func MiddlewareCtx(r *ghttp.Request) {
	service.Middleware().Ctx(r)
}

// MiddlewarePermissionAuth
// @Description: 访问权限鉴权
// @param r
func MiddlewarePermissionAuth(r *ghttp.Request) {
	service.Middleware().PermissionAuth(r)
}
