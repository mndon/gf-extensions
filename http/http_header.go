package http

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

const (
	HeaderRemoteIp = "X-Real-Ip"
	HeaderUA       = "User-Agent"
)

func GetRemoteIpFromCtx(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	return strings.Join(r.Header[HeaderRemoteIp], "; ")
}

func GetUaFromCtx(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	return strings.Join(r.Header[HeaderUA], "; ")
}
