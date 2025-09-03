/*
* @desc:token功能
* @company:
* @Author:
* @Date:   2022/9/27 17:01
 */

package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/tiger1103/gfast-token/gftoken"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		opt = getOptions()
		fun gftoken.OptionFunc
	)

	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func getOptions() *model.TokenOptions {
	ctx := gctx.New()
	var opt *model.TokenOptions
	v, _ := g.Cfg().Get(ctx, "admin.gfToken")
	if v != nil {
		v.Struct(&opt)
	} else {
		opt = &model.TokenOptions{
			CacheKey:     "admin:gfToken",
			Timeout:      10800,
			MaxRefresh:   5400,
			MultiLogin:   true,
			EncryptKey:   []byte("49c54195e750b04e74a8429b17896586"),
			ExcludePaths: []string{},
			CacheModel:   "memory",
		}
	}
	return opt
}

func init() {
	service.RegisterGToken(New())
}
