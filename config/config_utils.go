package config

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func GetValueFromConfigWithPanic(ctx context.Context, key string, def ...interface{}) *g.Var {
	value, err := g.Cfg().Get(ctx, key, def...)
	if err != nil {
		g.Log().Errorf(ctx, fmt.Sprintf("get %s from config error", key))
		panic(err)
	}
	return value
}
