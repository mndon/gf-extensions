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
)

func RouterRegister(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Login,
		)

		//登陆鉴权、用户信息注入、访问权限鉴权
		group.Middleware(MiddlewareLoginAuth, MiddlewareCtx, MiddlewarePermissionAuth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, OperationLog)
		// 绑定控制器
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
