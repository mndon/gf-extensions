package httpx

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const logFormat = `
=================[API_CALL]=================
uri : %s
url : %s
method : %s
body : %s
ip : %s
token : %s
userUid : %s
userAgent : %s 
%s=================[/API_CALL]=================
`

const errorLogFormat = "res: %s\nerr: %+v"

// MiddlewareLogRequest 验证请求签名
func MiddlewareLogRequest(r *ghttp.Request) {
	r.Middleware.Next()

	ctx := r.GetCtx()
	uri := r.Router.Uri
	url := r.URL.String()
	method := r.Method
	body := r.GetBodyString()
	ip := GetRemoteIpFromCtx(ctx)
	uId := GetIdentityFromCtx(ctx)
	token := r.Header.Get(HeaderAuthorization)
	ua := GetUaFromCtx(ctx)

	if err := r.GetError(); err != nil {
		logStr := fmt.Sprintf(logFormat, uri, url, method, body, ip, token, uId, ua, fmt.Sprintf(errorLogFormat, r.Response.BufferString(), err))
		g.Log().Warningf(ctx, logStr)
	} else {
		logStr := fmt.Sprintf(logFormat, uri, url, method, body, ip, token, uId, ua, "")
		g.Log().Info(ctx, logStr)
	}
}
