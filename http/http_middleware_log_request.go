package http

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

// MiddlewareLogRequest 验证请求签名
func MiddlewareLogRequest(r *ghttp.Request) {
	uId := gconv.String(JwtAuth().GetIdentity(r.GetCtx()))
	ua := strings.Join(r.Header[HeaderUA], "; ")
	ip := GetRemoteIpFromCtx(r.GetCtx())
	method := r.Method
	url := r.URL.String()
	uri := r.Router.Uri
	logStr := fmt.Sprintf("[MARK] [%s] [%s] [%s] [%s] [%s] [%s]", ip, method, url, uri, uId, ua)
	g.Log().Info(r.GetCtx(), logStr)
	r.Middleware.Next()
}
