package httpx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

const (
	HeaderForwardedFor  = "x-forwarded-for"
	HeaderRemoteIp      = "x-real-ip"
	HeaderXA            = "x-agent"
	HeaderUA            = "user-agent"
	HeaderAuthorization = "authorization"
)

func GetRemoteIpFromCtx(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r != nil {
		ip := r.Header.Get(HeaderForwardedFor)
		if ip == "" {
			ip = r.Header.Get(HeaderRemoteIp)
		}
		if idx := strings.Index(ip, ","); idx >= 0 {
			ip = ip[:idx]
		}
		return strings.TrimSpace(ip)
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
