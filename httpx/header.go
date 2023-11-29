package httpx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	HeaderRemoteIp      = "X-Real-Ip"
	HeaderXA            = "X-Agent"
	HeaderUA            = "User-Agent"
	HeaderAuthorization = "Authorization"
)

func GetRemoteIpFromCtx(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r != nil {
		return r.Header.Get(HeaderRemoteIp)
	}
	return ""
}

func GetUaFromCtx(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r != nil {
		return r.Header.Get(HeaderUA)
	}
	return ""
}
