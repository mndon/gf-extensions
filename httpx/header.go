package httpx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	HeaderRemoteIp      = "x-real-ip"
	HeaderXA            = "x-agent"
	HeaderUA            = "user-agent"
	HeaderAuthorization = "authorization"
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
