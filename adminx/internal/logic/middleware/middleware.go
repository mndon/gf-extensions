/*
* @desc:中间件处理
* @company:
* @Author:
* @Date:   2022/9/28 9:08
 */

package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		adminContext := new(model.Context)
		err = gconv.Struct(data.Data, &adminContext.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		service.Context().Init(r, adminContext)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// PermissionAuth 权限判断处理中间件
func (s *sMiddleware) PermissionAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	//获取登陆用户id
	adminId := service.Context().GetUserId(ctx)
	accessParams := r.Get("accessParams").Strings()
	accessParamsStr := ""
	if len(accessParams) > 0 && accessParams[0] != "undefined" {
		accessParamsStr = "?" + gstr.Join(accessParams, "&")
	}
	url := gstr.TrimLeft(r.Request.URL.Path, "/") + accessParamsStr
	//获取无需验证权限的用户id
	tagSuperAdmin := false
	service.AdminUser().NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == adminId {
			tagSuperAdmin = true
			return false
		}
		return true
	})
	if tagSuperAdmin {
		r.Middleware.Next()
		//不要再往后面执行
		return
	}
	//获取地址对应的菜单id
	menuList, err := service.AdminAuthRule().GetMenuList(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		r.SetError(gerror.WrapCode(gcode.New(5000, "请求数据失败", nil), err))
		return
	}
	var menu *model.AdminAuthRuleInfoRes
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			menu = m
			break
		}
	}
	//只验证存在数据库中的规则
	if menu != nil {
		//若是不登录能访问的接口则不判断权限
		excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
		for _, p := range excludePaths {
			if gstr.Equal(menu.Name, gstr.TrimLeft(p, "/")) {
				r.Middleware.Next()
				return
			}
		}
		//若存在不需要验证的条件则跳过
		if gstr.Equal(menu.Condition, "nocheck") {
			r.Middleware.Next()
			return
		}
		menuId := menu.Id
		//菜单没存数据库不验证权限
		if menuId != 0 {
			//判断权限操作
			enforcer, err := service.CasbinEnforcer(ctx)
			if err != nil {
				g.Log().Error(ctx, err)
				r.SetError(gerror.WrapCode(gcode.New(5000, "获取权限失败1", nil), err))
				return
			}
			hasAccess := false
			hasAccess, err = enforcer.Enforce(fmt.Sprintf("%s%d", service.AdminUser().GetCasBinUserPrefix(), adminId), gconv.String(menuId), "All")
			if err != nil {
				g.Log().Error(ctx, err)
				r.SetError(gerror.WrapCode(gcode.New(5000, "判断权限失败2", nil), err))
				return
			}
			if !hasAccess {
				r.SetError(gerror.WrapCode(gcode.New(4030, "没有访问权限1", nil), err))
				return
			}
		}
	} else if accessParamsStr != "" {
		r.SetError(gerror.WrapCode(gcode.New(4030, "没有访问权限2", nil), err))
		return
	}
	r.Middleware.Next()
}
