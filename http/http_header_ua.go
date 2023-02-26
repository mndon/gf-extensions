package http

import (
	"context"
	"regexp"
)

type Ua struct {
	ua string
}

func GetUaObjFromCtx(ctx context.Context) Ua {
	ua := GetUaFromCtx(ctx)
	return Ua{
		ua: ua,
	}
}

func (u *Ua) GetUa() string {
	return u.ua
}

func (u *Ua) GetChannelFormUA() string {
	return u.GetValueFormUA(`Channel\((.*?)\)`)
}

func (u *Ua) GetPlatformFormUA() string {
	return u.GetValueFormUA(`Platfom\((.*?)\)`)
}

func (u *Ua) GetAppVersionFormUA() string {
	return u.GetValueFormUA(`App Version\((.*?)\)`)
}

func (u *Ua) GetBrandFormUA() string {
	return u.GetValueFormUA(`Brand\((.*?)\)`)
}

func (u *Ua) GetModelFormUA() string {
	return u.GetValueFormUA(`Model\((.*?)\)`)
}

func (u *Ua) GetValueFormUA(matchKey string) string {
	brandReg := regexp.MustCompile(matchKey)
	result := brandReg.FindStringSubmatch(u.ua)
	if result != nil && len(result) == 2 {
		return result[1]
	}
	return ""
}
