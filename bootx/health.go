package bootx

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// enableHealthCheck
// @Description: 使能/health接口
// @param s
func enableHealthCheck(s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/health", healthCheck)
	})
}

func healthCheck(r *ghttp.Request) {
	ctx := r.Context()

	// 1. 检查数据库连接 (可选)
	// 如果你的应用强依赖数据库，建议加上这个检查

	if g.Cfg().MustGet(ctx, "health_check.db", false).Bool() {
		db := g.DB()
		err := db.PingMaster()
		if err != nil {
			g.Log().Error(ctx, "健康检查失败: 数据库连接异常", err)
			// 返回 503 状态码告知 Docker Swarm 容器目前不可用
			r.Response.WriteStatus(503, g.Map{
				"status":  "DOWN",
				"message": "Database connection failed",
			})
			return
		}
	}

	if g.Cfg().MustGet(ctx, "health_check.enable", false).Bool() {
		db := g.DB()
		err := db.PingMaster()
		if err != nil {
			_, err = g.Redis().Do(ctx, "PING")
			if err != nil {
				g.Log().Error(ctx, "健康检查失败: REDIS连接异常", err)
				// 返回 503 状态码告知 Docker Swarm 容器目前不可用
				r.Response.WriteStatus(503, g.Map{
					"status":  "DOWN",
					"message": "Redis connection failed",
				})
				return
			}
		}
	}

	// 3. 所有检查通过，返回 200 OK
	r.Response.WriteJson(g.Map{
		"status":  "UP",
		"version": "v1.0.0", // todo
	})
}
