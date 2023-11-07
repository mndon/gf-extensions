package httpx

import (
	"context"
	"regexp"
	"strings"
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

func (u *Ua) GetChannel() string {
	return u.GetValueFormUA(`Channel\((.*?)\)`)
}

func (u *Ua) GetPlatform() string {
	return u.GetValueFormUA(`Platfom\((.*?)\)`)
}

func (u *Ua) GetAppVersion() string {
	return u.GetValueFormUA(`App Version\((.*?)\)`)
}

func (u *Ua) GetBrand() string {
	return u.GetValueFormUA(`Brand\((.*?)\)`)
}

func (u *Ua) GetModel() string {
	return u.GetValueFormUA(`Model\((.*?)\)`)
}

func (u *Ua) GetPassport() string {
	return u.GetValueFormUA(`Passport\((.*?)\)`)
}

func (u *Ua) GetValueFormUA(matchKey string) string {
	brandReg := regexp.MustCompile(matchKey)
	result := brandReg.FindStringSubmatch(u.ua)
	if result != nil && len(result) == 2 {
		return strings.ToUpper(result[1])
	}
	return ""
}
