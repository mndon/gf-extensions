package httpx

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/sessionx"
	"github.com/mndon/gf-extensions/slicex"
	"time"
)

const (
	access     = "ACCESS"
	slowAccess = "SLOW_ACCESS"
)

type MiddlewareLogRequestBuilder struct {
	IgnoreUriMap        map[string]string //忽略打印请求map
	IgnoreReqDataUriMap map[string]string //忽略打印请求参数map
	SlowAccessThreshold int               //慢处理阈值，毫秒
}

func NewMiddlewareLogRequestBuilder() *MiddlewareLogRequestBuilder {
	return &MiddlewareLogRequestBuilder{
		IgnoreUriMap:        make(map[string]string),
		IgnoreReqDataUriMap: make(map[string]string),
		SlowAccessThreshold: 700,
	}
}

// SetSlowAccessThreshold
// @Description: 设置慢处理阈值
// @receiver m
// @param threshold
// @return *MiddlewareLogRequestBuilder
func (m *MiddlewareLogRequestBuilder) SetSlowAccessThreshold(threshold int) *MiddlewareLogRequestBuilder {
	m.SlowAccessThreshold = threshold
	return m
}

// SetIgnoreUri
// @Description: 设置忽略打印请求uri列表
// @receiver m
// @param uris
func (m *MiddlewareLogRequestBuilder) SetIgnoreUri(uris []string) *MiddlewareLogRequestBuilder {
	m.IgnoreUriMap = slicex.GroupByAndFlatten(uris, func(item string) string { return item })
	return m
}

// SetIgnoreReqDataUriMap
// @Description: 忽略打印请求参数uri列表
// @receiver m
// @param uris
func (m *MiddlewareLogRequestBuilder) SetIgnoreReqDataUriMap(uris []string) *MiddlewareLogRequestBuilder {
	m.IgnoreReqDataUriMap = slicex.GroupByAndFlatten(uris, func(item string) string { return item })
	return m
}

// Build
// @Description: 构建中间件
// @receiver m
// @param r
func (m *MiddlewareLogRequestBuilder) Build() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		r.Middleware.Next()

		// 命中忽略打印请求
		_, ok := m.IgnoreUriMap[r.RequestURI]
		if ok {
			return
		}

		ctx := r.GetCtx()
		mark := access
		accessTime := int(gtime.Now().Sub(r.EnterTime) / time.Millisecond)
		if accessTime > m.SlowAccessThreshold {
			mark = slowAccess
		}

		reqBody := r.GetBodyString()
		// 命中忽略打印请求体
		_, ok = m.IgnoreReqDataUriMap[r.RequestURI]
		if ok {
			reqBody = ""
		}

		content := fmt.Sprintf(
			`[%d ms] [%s] %d "%s %s %s", "%s", "%s", "%s", "%s", "%s"`,
			accessTime, mark, r.Response.Status, r.Method, r.Router.Uri, r.URL.String(), reqBody,
			GetRemoteIpFromCtx(ctx), sessionx.GetUserUid(ctx), GetAgent(ctx).agent, r.Header.Get(HeaderAuthorization),
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
}
