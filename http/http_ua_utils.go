package http

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"regexp"
	"strings"
)

type UaUtil struct {
	ua  string
	ctx context.Context
}

func NewUaUtil(ctx context.Context) UaUtil {
	return UaUtil{
		ctx: ctx,
	}
}

func (u *UaUtil) GetChannelFormUA() string {
	return u.GetValueFormUA(`Channel\((.*?)\)`)
}

func (u *UaUtil) GetPlatformFormUA() string {
	return u.GetValueFormUA(`Platfom\((.*?)\)`)
}

func (u *UaUtil) GetAppVersionFormUA() string {
	return u.GetValueFormUA(`App Version\((.*?)\)`)
}

func (u *UaUtil) GetBrandFormUA() string {
	return u.GetValueFormUA(`Brand\((.*?)\)`)
}

func (u *UaUtil) GetValueFormUA(matchKey string) string {
	brandReg := regexp.MustCompile(matchKey)
	result := brandReg.FindStringSubmatch(u.GetUaFromContext())
	if result != nil && len(result) == 2 {
		return result[1]
	}
	return ""
}

func (u *UaUtil) GetUaFromContext() string {
	if u.ua == "" {
		r := ghttp.RequestFromCtx(u.ctx)
		u.ua = strings.Join(r.Header["User-Agent"], "; ")
	}
	return u.ua
}
