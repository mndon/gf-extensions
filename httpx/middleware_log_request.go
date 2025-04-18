package httpx

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/sessionx"
	"time"
)

const (
	access     = "ACCESS"
	slowAccess = "SLOW_ACCESS"
)

var accessLogger *glog.Logger

func init() {
	config := g.Log().GetConfig()
	config.StStatus = 0
	accessLogger = glog.New()
	accessLogger.SetConfig(config)
}

// MiddlewareLogRequest 验证请求签名
func MiddlewareLogRequest(r *ghttp.Request) {
	r.Middleware.Next()

	ctx := r.GetCtx()
	mark := access
	accessTime := gtime.Now().Sub(r.EnterTime) / time.Millisecond
	if accessTime > 700 {
		mark = slowAccess
	}

	bodyString := r.GetBodyString()
	if len(bodyString) > 512 {
		bodyString = bodyString[:512] + "..."
	}

	content := fmt.Sprintf(
		`[%d ms] [%s] %d "%s %s %s", "%s", "%s", "%s", "%s", "%s"`,
		accessTime, mark, r.Response.Status, r.Method, r.Router.Uri, r.URL.String(), bodyString,
		GetRemoteIpFromCtx(ctx), sessionx.GetUserUid(ctx), AgentStrFromHeader(ctx), r.Header.Get(HeaderAuthorization),
	)

	err := r.GetError()
	if err != nil {
		content += fmt.Sprintf(`, "%s"`, r.Response.BufferString())
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += ", " + err.Error()
		}

		code := gerror.Code(err)
		if code != gcode.CodeNil {
			accessLogger.Warning(ctx, content)
		} else {
			accessLogger.Error(ctx, content)
		}
	} else {
		accessLogger.Info(ctx, content)
	}
}
