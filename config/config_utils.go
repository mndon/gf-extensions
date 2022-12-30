package config

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func GetWithCmdFromCfgWithPanic(ctx context.Context, key string, def ...interface{}) *g.Var {
	value, err := g.Cfg().GetWithCmd(ctx, key, def...)
	if err != nil {
		g.Log().Errorf(ctx, fmt.Sprintf("get %s from config error", key))
		panic(err)
	}
	return value
}
