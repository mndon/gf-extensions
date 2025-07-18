/*
* @desc:登录
* @company:
* @Author:
* @Date:   2022/4/27 21:52
 */

package controller

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/lib/libUtils"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var (
	Login = loginController{}
)

type loginController struct {
	BaseController
}

func (c *loginController) Login(ctx context.Context, req *api.UserLoginReq) (res *api.UserLoginRes, err error) {
	var (
		user        *model.LoginUserRes
		token       string
		permissions []string
		menuList    []*model.UserMenus
	)
	//判断验证码是否正确
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.AdminUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		// 保存登录失败的日志信息
		service.AdminLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
			Status:    0,
			Username:  req.Username,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		return
	}
	err = service.AdminUser().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 报存登录成功的日志信息
	service.AdminLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
		Status:    1,
		Username:  req.Username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	//获取用户菜单数据
	menuList, permissions, err = service.AdminUser().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	res = &api.UserLoginRes{
		UserInfo:    user,
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
	}
	//用户在线状态保存
	service.AdminUserOnline().Invoke(gctx.New(), &model.AdminUserOnlineParams{
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(token),
		Token:     token,
		Username:  user.UserName,
		Ip:        ip,
	})
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *api.UserLoginOutReq) (res *api.UserLoginOutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
