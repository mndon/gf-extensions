/*
* @desc:缓存处理
* @company:
* @Author:
* @Date:   2023/2/1 18:14
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/service"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Cache = new(cacheController)

type cacheController struct {
	BaseController
}

func (c *cacheController) Remove(ctx context.Context, req *api.CacheRemoveReq) (res *api.CacheRemoveRes, err error) {
	service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
	service.Cache().RemoveByTag(ctx, consts.CacheAdminConfigTag)
	service.Cache().RemoveByTag(ctx, consts.CacheAdminAuthTag)
	cacheRedis := g.Cfg().MustGet(ctx, "api.cache.model").String()
	if cacheRedis == consts.CacheModelRedis {
		cursor := 0
		cachePrefix := g.Cfg().MustGet(ctx, "api.cache.prefix").String()
		cachePrefix += consts.CachePrefix
		for {
			var v *gvar.Var
			v, err = g.Redis().Do(ctx, "scan", cursor, "match", cachePrefix+"*", "count", "100")
			if err != nil {
				return
			}
			data := gconv.SliceAny(v)
			var dataSlice []string
			err = gconv.Structs(data[1], &dataSlice)
			if err != nil {
				return
			}
			for _, d := range dataSlice {
				_, err = g.Redis().Do(ctx, "del", d)
				if err != nil {
					return
				}
			}
			cursor = gconv.Int(data[0])
			if cursor == 0 {
				break
			}
		}
	}
	return
}
