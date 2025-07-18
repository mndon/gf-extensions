/*
* @desc:系统参数配置
* @company:
* @Author:
* @Date:   2022/4/18 21:17
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var Config = configController{}

type configController struct {
	BaseController
}

// List 系统参数列表
func (c *configController) List(ctx context.Context, req *api.ConfigSearchReq) (res *api.ConfigSearchRes, err error) {
	res, err = service.AdminConfig().List(ctx, req)
	return
}

// Add 添加系统参数
func (c *configController) Add(ctx context.Context, req *api.ConfigAddReq) (res *api.ConfigAddRes, err error) {
	err = service.AdminConfig().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取系统参数
func (c *configController) Get(ctx context.Context, req *api.ConfigGetReq) (res *api.ConfigGetRes, err error) {
	res, err = service.AdminConfig().Get(ctx, req.Id)
	return
}

// Edit 修改系统参数
func (c *configController) Edit(ctx context.Context, req *api.ConfigEditReq) (res *api.ConfigEditRes, err error) {
	err = service.AdminConfig().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Delete 删除系统参数
func (c *configController) Delete(ctx context.Context, req *api.ConfigDeleteReq) (res *api.ConfigDeleteRes, err error) {
	err = service.AdminConfig().Delete(ctx, req.Ids)
	return
}
