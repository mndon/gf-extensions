/*
* @desc:在线用户管理
* @company:
* @Author:
* @Date:   2023/1/10 17:23
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var UserOnline = new(AdminUserOnlineController)

type AdminUserOnlineController struct{}

func (c *AdminUserOnlineController) List(ctx context.Context, req *api.AdminUserOnlineSearchReq) (res *api.AdminUserOnlineSearchRes, err error) {
	res, err = service.AdminUserOnline().GetOnlineListPage(ctx, req)
	return
}

func (c *AdminUserOnlineController) ForceLogout(ctx context.Context, req *api.AdminUserOnlineForceLogoutReq) (res *api.AdminUserOnlineForceLogoutRes, err error) {
	err = service.AdminUserOnline().ForceLogout(ctx, req.Ids)
	return
}
