/*
* @desc:登录日志管理
* @company:
* @Author:
* @Date:   2022/4/24 22:14
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var LoginLog = loginLogController{}

type loginLogController struct {
	BaseController
}

func (c *loginLogController) List(ctx context.Context, req *api.LoginLogSearchReq) (res *api.LoginLogSearchRes, err error) {
	res, err = service.AdminLoginLog().List(ctx, req)
	return
}

func (c *loginLogController) Delete(ctx context.Context, req *api.LoginLogDelReq) (res *api.LoginLogDelRes, err error) {
	err = service.AdminLoginLog().DeleteLoginLogByIds(ctx, req.Ids)
	return
}

func (c *loginLogController) Clear(ctx context.Context, req *api.LoginLogClearReq) (res *api.LoginLogClearRes, err error) {
	err = service.AdminLoginLog().ClearLoginLog(ctx)
	return
}
