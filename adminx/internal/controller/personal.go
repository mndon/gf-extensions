/*
* @desc:xxxx功能描述
* @company:
* @Author:
* @Date:   2022/11/3 10:32
 */

package controller

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/lib/libUtils"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var Personal = new(personalController)

type personalController struct {
}

func (c *personalController) GetPersonal(ctx context.Context, req *api.PersonalInfoReq) (res *api.PersonalInfoRes, err error) {
	res, err = service.Personal().GetPersonalInfo(ctx, req)
	return
}

func (c *personalController) EditPersonal(ctx context.Context, req *api.PersonalEditReq) (res *api.PersonalEditRes, err error) {
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	res = new(api.PersonalEditRes)
	res.UserInfo, err = service.Personal().EditPersonal(ctx, req)
	if err != nil {
		return
	}
	key := gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword+ip+userAgent)
	}
	res.UserInfo.UserPassword = ""
	res.Token, err = service.GfToken().GenerateToken(ctx, key, res.UserInfo)
	return
}

func (c *personalController) ResetPwdPersonal(ctx context.Context, req *api.PersonalResetPwdReq) (res *api.PersonalResetPwdRes, err error) {
	res, err = service.Personal().ResetPwdPersonal(ctx, req)
	return
}
