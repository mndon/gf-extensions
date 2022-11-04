package http

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddlewareAuth struct{}

var insMiddlewareAuth = sMiddlewareAuth{}

func MiddlewareAuth() *sMiddlewareAuth {
	return &insMiddlewareAuth
}

func (s *sMiddlewareAuth) JwtAuth(r *ghttp.Request) {
	JwtAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
