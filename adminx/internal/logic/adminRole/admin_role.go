/*
* @desc:角色管理
* @company:
* @Author:
* @Date:   2022/9/26 15:54
 */

package adminRole

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterAdminRole(New())
}

func New() *sAdminRole {
	return &sAdminRole{}
}

type sAdminRole struct {
}

func (s *sAdminRole) GetRoleListSearch(ctx context.Context, req *api.RoleListReq) (res *api.RoleListRes, err error) {
	res = new(api.RoleListRes)
	g.Try(ctx, func(ctx context.Context) {
		model := dao.AdminRole.Ctx(ctx)
		if req.RoleName != "" {
			model = model.Where("name like ?", "%"+req.RoleName+"%")
		}
		if req.Status != "" {
			model = model.Where("status", gconv.Int(req.Status))
		}
		res.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取角色数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		err = model.Page(res.CurrentPage, req.PageSize).Order("id asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

// GetRoleList 获取角色列表
func (s *sAdminRole) GetRoleList(ctx context.Context) (list []*entity.AdminRole, err error) {
	cache := service.Cache()
	//从缓存获取
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheAdminRole, s.getRoleListFromDb, 0, consts.CacheAdminAuthTag)
	if !iList.IsEmpty() {
		err = gconv.Struct(iList, &list)
	}
	return
}

// 从数据库获取所有角色
func (s *sAdminRole) getRoleListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*entity.AdminRole
		//从数据库获取
		err = dao.AdminRole.Ctx(ctx).
			Order(dao.AdminRole.Columns().ListOrder + " asc," + dao.AdminRole.Columns().Id + " asc").
			Scan(&v)
		liberr.ErrIsNil(ctx, err, "获取角色数据失败")
		value = v
	})
	return
}

// AddRoleRule 添加角色权限
func (s *sAdminRole) AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		ruleIdsStr := gconv.Strings(ruleIds)
		for _, v := range ruleIdsStr {
			_, err = enforcer.AddPolicy(gconv.String(roleId), v, "All")
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// DelRoleRule 删除角色权限
func (s *sAdminRole) DelRoleRule(ctx context.Context, roleId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		_, err = enforcer.RemoveFilteredPolicy(0, gconv.String(roleId))
		liberr.ErrIsNil(ctx, e)
	})
	return
}

func (s *sAdminRole) AddRole(ctx context.Context, req *api.RoleAddReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			roleId, e := dao.AdminRole.Ctx(ctx).TX(tx).InsertAndGetId(req)
			liberr.ErrIsNil(ctx, e, "添加角色失败")
			//添加角色权限
			e = s.AddRoleRule(ctx, req.MenuIds, roleId)
			liberr.ErrIsNil(ctx, e)
			//清除缓存
			service.Cache().Remove(ctx, consts.CacheAdminRole)
		})
		return err
	})
	return
}

func (s *sAdminRole) Get(ctx context.Context, id uint) (res *entity.AdminRole, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminRole.Ctx(ctx).WherePri(id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取角色信息失败")
	})
	return
}

// GetFilteredNamedPolicy 获取角色关联的菜单规则
func (s *sAdminRole) GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := service.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		gp := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
		gpSlice = make([]int, len(gp))
		for k, v := range gp {
			gpSlice[k] = gconv.Int(v[1])
		}
	})
	return
}

// EditRole 修改角色
func (s *sAdminRole) EditRole(ctx context.Context, req *api.RoleEditReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.AdminRole.Ctx(ctx).TX(tx).WherePri(req.Id).Data(&do.AdminRole{
				Status:    req.Status,
				ListOrder: req.ListOrder,
				Name:      req.Name,
				Remark:    req.Remark,
			}).Update()
			liberr.ErrIsNil(ctx, e, "修改角色失败")
			//删除角色权限
			e = s.DelRoleRule(ctx, req.Id)
			liberr.ErrIsNil(ctx, e)
			//添加角色权限
			e = s.AddRoleRule(ctx, req.MenuIds, req.Id)
			liberr.ErrIsNil(ctx, e)
			//清除缓存
			service.Cache().Remove(ctx, consts.CacheAdminRole)
		})
		return err
	})
	return
}

// DeleteByIds 删除角色
func (s *sAdminRole) DeleteByIds(ctx context.Context, ids []int64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AdminRole.Ctx(ctx).TX(tx).Where(dao.AdminRole.Columns().Id+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除角色失败")
			//删除角色权限
			for _, v := range ids {
				err = s.DelRoleRule(ctx, v)
				liberr.ErrIsNil(ctx, err)
			}
			//清除缓存
			service.Cache().Remove(ctx, consts.CacheAdminRole)
		})
		return err
	})
	return
}
