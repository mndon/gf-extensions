package adminx

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func MiddlewareAdminAuth(group *ghttp.RouterGroup) {
	_ = service.GfToken().Middleware(group)
}
