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
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/tiger1103/gfast-token/gftoken"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	liberr.ErrIsNil(ctx, err)
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

func init() {
	service.RegisterGToken(New())
}
