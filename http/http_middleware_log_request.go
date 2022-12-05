package http

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

const logFormat = `
=================[API_CALL]=================
uri : %s
url : %s
method : %s
body : %s
ip : %s
token : %s
userId : %s
userAgent : %s 
=================[/API_CALL]=================
`

// MiddlewareLogRequest 验证请求签名
func MiddlewareLogRequest(r *ghttp.Request) {
	ctx := r.GetCtx()
	uri := r.Router.Uri
	url := r.URL.String()
	method := r.Method
	body := r.GetBodyString()
	ip := GetRemoteIpFromCtx(ctx)
	uId := gconv.String(JwtAuth().GetIdentity(ctx))
	token := r.Header.Get(HeaderAuthorization)
	ua := GetUaFromCtx(ctx)

	logStr := fmt.Sprintf(logFormat, uri, url, method, body, ip, token, uId, ua)
	g.Log().Info(ctx, logStr)
	r.Middleware.Next()
}
