package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var (
	User = userController{}
)

type userController struct {
	BaseController
}

// GetUserMenus 获取用户菜单及按钮权限
func (c *userController) GetUserMenus(ctx context.Context, req *api.UserMenusReq) (res *api.UserMenusRes, err error) {
	var (
		permissions []string
		menuList    []*model.UserMenus
	)
	userId := service.Context().GetUserId(ctx)
	menuList, permissions, err = service.AdminUser().GetAdminRules(ctx, userId)
	res = &api.UserMenusRes{
		MenuList:    menuList,
		Permissions: permissions,
	}
	return
}

// List 用户列表
func (c *userController) List(ctx context.Context, req *api.UserSearchReq) (res *api.UserSearchRes, err error) {
	var (
		total    interface{}
		userList []*entity.AdminUser
	)
	res = new(api.UserSearchRes)
	total, userList, err = service.AdminUser().List(ctx, req)
	if err != nil || total == 0 {
		return
	}
	res.Total = total
	res.UserList, err = service.AdminUser().GetUsersRoleDept(ctx, userList)
	return
}

// GetParams 获取用户维护相关参数
func (c *userController) GetParams(ctx context.Context, req *api.UserGetParamsReq) (res *api.UserGetParamsRes, err error) {
	res = new(api.UserGetParamsRes)
	res.RoleList, err = service.AdminRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	return
}

// Add 添加用户
func (c *userController) Add(ctx context.Context, req *api.UserAddReq) (res *api.UserAddRes, err error) {
	err = service.AdminUser().Add(ctx, req)
	return
}

// GetEditUser 获取修改用户信息
func (c *userController) GetEditUser(ctx context.Context, req *api.UserGetEditReq) (res *api.UserGetEditRes, err error) {
	res, err = service.AdminUser().GetEditUser(ctx, req.Id)
	return
}

// Edit 修改用户
func (c *userController) Edit(ctx context.Context, req *api.UserEditReq) (res *api.UserEditRes, err error) {
	err = service.AdminUser().Edit(ctx, req)
	return
}

// ResetPwd 重置密码
func (c *userController) ResetPwd(ctx context.Context, req *api.UserResetPwdReq) (res *api.UserResetPwdRes, err error) {
	err = service.AdminUser().ResetUserPwd(ctx, req)
	return
}

// SetStatus 修改用户状态
func (c *userController) SetStatus(ctx context.Context, req *api.UserStatusReq) (res *api.UserStatusRes, err error) {
	err = service.AdminUser().ChangeUserStatus(ctx, req)
	return
}

// Delete 删除用户
func (c *userController) Delete(ctx context.Context, req *api.UserDeleteReq) (res *api.UserDeleteRes, err error) {
	err = service.AdminUser().Delete(ctx, req.Ids)
	return
}

// GetUsers 通过用户id批量获取用户信息
func (c *userController) GetUsers(ctx context.Context, req *api.UserGetByIdsReq) (res *api.UserGetByIdsRes, err error) {
	res = new(api.UserGetByIdsRes)
	res.List, err = service.AdminUser().GetUsers(ctx, req.Ids)
	return
}
