/*
* @desc:菜单处理
* @company:
* @Author:
* @Date:   2022/9/23 16:14
 */

package adminAuthRule

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterAdminAuthRule(New())
}

func New() *sAdminAuthRule {
	return &sAdminAuthRule{}
}

type sAdminAuthRule struct {
}

func (s *sAdminAuthRule) GetMenuListSearch(ctx context.Context, req *api.RuleSearchReq) (res []*model.AdminAuthRuleInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminAuthRule.Ctx(ctx)
		if req.Title != "" {
			m = m.Where("title like ?", "%"+req.Title+"%")
		}
		if req.Component != "" {
			m = m.Where("component like ?", "%"+req.Component+"%")
		}
		err = m.Fields(model.AdminAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取菜单失败")
	})
	return
}

// GetIsMenuList 获取isMenu=0|1
func (s *sAdminAuthRule) GetIsMenuList(ctx context.Context) ([]*model.AdminAuthRuleInfoRes, error) {
	list, err := s.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	var gList = make([]*model.AdminAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.MenuType == 0 || v.MenuType == 1 {
			gList = append(gList, v)
		}
	}
	return gList, nil
}

// GetMenuList 获取所有菜单
func (s *sAdminAuthRule) GetMenuList(ctx context.Context) (list []*model.AdminAuthRuleInfoRes, err error) {
	cache := service.Cache()
	//从缓存获取
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheAdminAuthMenu, s.getMenuListFromDb, 0, consts.CacheAdminAuthTag)
	if !iList.IsEmpty() {
		err = gconv.Struct(iList, &list)
		liberr.ErrIsNil(ctx, err)
	}
	return
}

// 从数据库获取所有菜单
func (s *sAdminAuthRule) getMenuListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*model.AdminAuthRuleInfoRes
		//从数据库获取
		err = dao.AdminAuthRule.Ctx(ctx).
			Fields(model.AdminAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&v)
		liberr.ErrIsNil(ctx, err, "获取菜单数据失败")
		value = v
	})
	return
}

// GetIsButtonList 获取所有按钮isMenu=2 菜单列表
func (s *sAdminAuthRule) GetIsButtonList(ctx context.Context) ([]*model.AdminAuthRuleInfoRes, error) {
	list, err := s.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	var gList = make([]*model.AdminAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.MenuType == 2 {
			gList = append(gList, v)
		}
	}
	return gList, nil
}

// Add 添加菜单
func (s *sAdminAuthRule) Add(ctx context.Context, req *api.RuleAddReq) (err error) {
	if s.menuNameExists(ctx, req.Name, 0) {
		err = gerror.New("接口规则已经存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//菜单数据
			data := do.AdminAuthRule{
				Pid:       req.Pid,
				Name:      req.Name,
				Title:     req.Title,
				Icon:      req.Icon,
				Condition: req.Condition,
				Remark:    req.Remark,
				MenuType:  req.MenuType,
				Weigh:     req.Weigh,
				IsHide:    req.IsHide,
				Path:      req.Path,
				Component: req.Component,
				IsLink:    req.IsLink,
				IsIframe:  req.IsIframe,
				IsCached:  req.IsCached,
				Redirect:  req.Redirect,
				IsAffix:   req.IsAffix,
				LinkUrl:   req.LinkUrl,
			}
			ruleId, e := dao.AdminAuthRule.Ctx(ctx).TX(tx).InsertAndGetId(data)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
			e = s.BindRoleRule(ctx, ruleId, req.Roles)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
		})
		return err
	})
	if err == nil {
		// 删除相关缓存
		service.Cache().Remove(ctx, consts.CacheAdminAuthMenu)
	}
	return
}

// 检查菜单规则是否存在
func (s *sAdminAuthRule) menuNameExists(ctx context.Context, name string, id uint) bool {
	m := dao.AdminAuthRule.Ctx(ctx).Where("name=?", name)
	if id != 0 {
		m = m.Where("id!=?", id)
	}
	c, err := m.Fields(dao.AdminAuthRule.Columns().Id).Limit(1).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}
	return !c.IsEmpty()
}

// BindRoleRule 绑定角色权限
func (s *sAdminAuthRule) BindRoleRule(ctx context.Context, ruleId interface{}, roleIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		for _, roleId := range roleIds {
			_, err = enforcer.AddPolicy(fmt.Sprintf("%d", roleId), fmt.Sprintf("%d", ruleId), "All")
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sAdminAuthRule) Get(ctx context.Context, id uint) (rule *entity.AdminAuthRule, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminAuthRule.Ctx(ctx).WherePri(id).Scan(&rule)
		liberr.ErrIsNil(ctx, err, "获取菜单失败")
	})
	return
}

func (s *sAdminAuthRule) GetMenuRoles(ctx context.Context, id uint) (roleIds []uint, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		policies := enforcer.GetFilteredNamedPolicy("p", 1, gconv.String(id))
		for _, policy := range policies {
			roleIds = append(roleIds, gconv.Uint(policy[0]))
		}
	})
	return
}

func (s *sAdminAuthRule) Update(ctx context.Context, req *api.RuleUpdateReq) (err error) {
	if s.menuNameExists(ctx, req.Name, req.Id) {
		err = gerror.New("接口规则已经存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//菜单数据
			data := do.AdminAuthRule{
				Pid:       req.Pid,
				Name:      req.Name,
				Title:     req.Title,
				Icon:      req.Icon,
				Condition: req.Condition,
				Remark:    req.Remark,
				MenuType:  req.MenuType,
				Weigh:     req.Weigh,
				IsHide:    req.IsHide,
				Path:      req.Path,
				Component: req.Component,
				IsLink:    req.IsLink,
				IsIframe:  req.IsIframe,
				IsCached:  req.IsCached,
				Redirect:  req.Redirect,
				IsAffix:   req.IsAffix,
				LinkUrl:   req.LinkUrl,
			}
			_, e := dao.AdminAuthRule.Ctx(ctx).TX(tx).WherePri(req.Id).Update(data)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
			e = s.UpdateRoleRule(ctx, req.Id, req.Roles)
			liberr.ErrIsNil(ctx, e, "添加菜单失败")
		})
		return err
	})
	if err == nil {
		// 删除相关缓存
		service.Cache().Remove(ctx, consts.CacheAdminAuthMenu)
	}
	return
}

func (s *sAdminAuthRule) UpdateRoleRule(ctx context.Context, ruleId uint, roleIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		//删除旧权限
		_, e = enforcer.RemoveFilteredPolicy(1, gconv.String(ruleId))
		liberr.ErrIsNil(ctx, e)
		// 添加新权限
		roleIdsStrArr := gconv.Strings(roleIds)
		for _, v := range roleIdsStrArr {
			_, e = enforcer.AddPolicy(v, gconv.String(ruleId), "All")
			liberr.ErrIsNil(ctx, e)
		}
	})
	return
}

func (s *sAdminAuthRule) GetMenuListTree(pid uint, list []*model.AdminAuthRuleInfoRes) []*model.AdminAuthRuleTreeRes {
	tree := make([]*model.AdminAuthRuleTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == pid {
			t := &model.AdminAuthRuleTreeRes{
				AdminAuthRuleInfoRes: menu,
			}
			child := s.GetMenuListTree(menu.Id, list)
			if child != nil {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}

// DeleteMenuByIds 删除菜单
func (s *sAdminAuthRule) DeleteMenuByIds(ctx context.Context, ids []int) (err error) {
	var list []*model.AdminAuthRuleInfoRes
	list, err = s.GetMenuList(ctx)
	if err != nil {
		return
	}
	childrenIds := make([]int, 0, len(list))
	for _, id := range ids {
		rules := s.FindSonByParentId(list, gconv.Uint(id))
		for _, child := range rules {
			childrenIds = append(childrenIds, gconv.Int(child.Id))
		}
	}
	ids = append(ids, childrenIds...)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AdminAuthRule.Ctx(ctx).Where("id in (?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除失败")
			//删除权限
			enforcer, err := service.CasbinEnforcer(ctx)
			liberr.ErrIsNil(ctx, err)
			for _, v := range ids {
				_, err = enforcer.RemoveFilteredPolicy(1, gconv.String(v))
				liberr.ErrIsNil(ctx, err)
			}
			// 删除相关缓存
			service.Cache().Remove(ctx, consts.CacheAdminAuthMenu)
		})
	})
	return
}

func (s *sAdminAuthRule) FindSonByParentId(list []*model.AdminAuthRuleInfoRes, pid uint) []*model.AdminAuthRuleInfoRes {
	children := make([]*model.AdminAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.Pid == pid {
			children = append(children, v)
			fChildren := s.FindSonByParentId(list, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}
