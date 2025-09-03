/*
* @desc:缓存处理
* @company:
* @Author:
* @Date:   2022/9/27 16:33
 */

package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/tiger1103/gfast-cache/cache"
)

func init() {
	service.RegisterCache(New())
}

func New() *sCache {
	var (
		ctx            = gctx.New()
		cacheContainer *cache.GfCache
	)
	prefix := g.Cfg().MustGet(ctx, "admin.cache.prefix").String()
	model := g.Cfg().MustGet(ctx, "admin.cache.model").String()
	if model == consts.CacheModelRedis {
		// redis
		cacheContainer = cache.NewRedis(prefix)
	} else {
		// memory
		cacheContainer = cache.New(prefix)
	}
	return &sCache{
		GfCache: cacheContainer,
		prefix:  prefix,
	}
}

type sCache struct {
	*cache.GfCache
	prefix string
}
