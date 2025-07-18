/*
* @desc:菜单
* @company:
* @Author:
* @Date:   2022/3/16 10:36
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var Menu = menuController{}

type menuController struct {
	BaseController
}

func (c *menuController) List(ctx context.Context, req *api.RuleSearchReq) (res *api.RuleListRes, err error) {
	var list []*model.AdminAuthRuleInfoRes
	res = &api.RuleListRes{
		Rules: make([]*model.AdminAuthRuleTreeRes, 0),
	}
	list, err = service.AdminAuthRule().GetMenuListSearch(ctx, req)
	if req.Title != "" || req.Component != "" {
		for _, menu := range list {
			res.Rules = append(res.Rules, &model.AdminAuthRuleTreeRes{
				AdminAuthRuleInfoRes: menu,
			})
		}
	} else {
		res.Rules = service.AdminAuthRule().GetMenuListTree(0, list)
	}
	return
}

func (c *menuController) Add(ctx context.Context, req *api.RuleAddReq) (res *api.RuleAddRes, err error) {
	err = service.AdminAuthRule().Add(ctx, req)
	return
}

// GetAddParams 获取菜单添加及编辑相关参数
func (c *menuController) GetAddParams(ctx context.Context, req *api.RuleGetParamsReq) (res *api.RuleGetParamsRes, err error) {
	// 获取角色列表
	res = new(api.RuleGetParamsRes)
	res.Roles, err = service.AdminRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Menus, err = service.AdminAuthRule().GetIsMenuList(ctx)
	return
}

// Get 获取菜单信息
func (c *menuController) Get(ctx context.Context, req *api.RuleInfoReq) (res *api.RuleInfoRes, err error) {
	res = new(api.RuleInfoRes)
	res.Rule, err = service.AdminAuthRule().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.RoleIds, err = service.AdminAuthRule().GetMenuRoles(ctx, req.Id)
	return
}

// Update 菜单修改
func (c *menuController) Update(ctx context.Context, req *api.RuleUpdateReq) (res *api.RuleUpdateRes, err error) {
	err = service.AdminAuthRule().Update(ctx, req)
	return
}

// Delete 删除菜单
func (c *menuController) Delete(ctx context.Context, req *api.RuleDeleteReq) (res *api.RuleDeleteRes, err error) {
	err = service.AdminAuthRule().DeleteMenuByIds(ctx, req.Ids)
	return
}
