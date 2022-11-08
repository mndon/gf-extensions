package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareJwtAuth(r *ghttp.Request) {
	JwtAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
