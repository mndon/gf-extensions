package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func JwtAuthMiddleware(r *ghttp.Request) {
	JwtAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
