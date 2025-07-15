/*
* @desc:后台路由
* @company:
* @Author:
* @Date:   2022/2/18 17:34
 */

package adminx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/adminx/internal/controller"
	_ "github.com/mndon/gf-extensions/adminx/internal/logic"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func RouterRegister(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Login,
		)
		//登录验证拦截
		MiddlewareAdminAuth(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		group.Bind(
			controller.User,
			controller.Menu,
			controller.Role,
			controller.DictType,
			controller.DictData,
			controller.Config,
			controller.Monitor,
			controller.LoginLog,
			controller.OperLog,
			controller.Personal,
			controller.UserOnline,
			controller.Cache, // 缓存处理
		)
	})
}
