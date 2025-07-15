/*
* @desc:角色管理
* @company:
* @Author:
* @Date:   2022/3/30 9:08
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var Role = roleController{}

type roleController struct {
	BaseController
}

// List 角色列表
func (c *roleController) List(ctx context.Context, req *api.RoleListReq) (res *api.RoleListRes, err error) {
	res, err = service.AdminRole().GetRoleListSearch(ctx, req)
	return
}

// GetParams 获取角色表单参数
func (c *roleController) GetParams(ctx context.Context, req *api.RoleGetParamsReq) (res *api.RoleGetParamsRes, err error) {
	res = new(api.RoleGetParamsRes)
	res.Menu, err = service.AdminAuthRule().GetMenuList(ctx)
	return
}

// Add 添加角色信息
func (c *roleController) Add(ctx context.Context, req *api.RoleAddReq) (res *api.RoleAddRes, err error) {
	err = service.AdminRole().AddRole(ctx, req)
	return
}

// Get 获取角色信息
func (c *roleController) Get(ctx context.Context, req *api.RoleGetReq) (res *api.RoleGetRes, err error) {
	res = new(api.RoleGetRes)
	res.Role, err = service.AdminRole().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.AdminRole().GetFilteredNamedPolicy(ctx, req.Id)
	return
}

// Edit 修改角色信息
func (c *roleController) Edit(ctx context.Context, req *api.RoleEditReq) (res *api.RoleEditRes, err error) {
	err = service.AdminRole().EditRole(ctx, req)
	return
}

// Delete 删除角色
func (c *roleController) Delete(ctx context.Context, req *api.RoleDeleteReq) (res *api.RoleDeleteRes, err error) {
	err = service.AdminRole().DeleteByIds(ctx, req.Ids)
	return
}
