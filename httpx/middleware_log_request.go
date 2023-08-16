package httpx

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/sessionx"
)

// MiddlewareLogRequest 验证请求签名
func MiddlewareLogRequest(r *ghttp.Request) {
	r.Middleware.Next()

	ctx := r.GetCtx()
	content := fmt.Sprintf(
		`[ACCESS] %d "%s %s %s" "%s" %.3f, "%s", "%s", "%s", "%s"`,
		r.Response.Status, r.Method, r.Router.Uri, r.URL.String(), r.GetBodyString(), float64(gtime.TimestampMilli()-r.EnterTime)/1000,
		GetRemoteIpFromCtx(ctx), sessionx.GetUserUid(ctx), r.UserAgent(), r.Header.Get(HeaderAuthorization),
	)

	if err := r.GetError(); err != nil {
		content += fmt.Sprintf(`, "%s"`, r.Response.BufferString())
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += ", " + err.Error()
		}
	}
	g.Log().Print(ctx, content)
}
