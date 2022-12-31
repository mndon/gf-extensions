package config

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

// GetWithCmdFromCfgWithPanic
// @Description: 获取配置， 先从cmd获取， 获取不到再到配置文件获取
// @param ctx
// @param key
// @param def
// @return *g.Var
func GetWithCmdFromCfgWithPanic(ctx context.Context, key string, def ...interface{}) *g.Var {
	// 先从cmd获取
	if v := gcmd.GetOpt(key); v != nil {
		return gvar.New(v)
	}
	// 再从配置文件获取
	value, err := g.Cfg().Get(ctx, key)
	if err != nil {
		g.Log().Errorf(ctx, fmt.Sprintf("get %s from config error", key))
		panic(err)
	}
	if value == nil {
		if len(def) > 0 {
			return gvar.New(def[0])
		}
		return nil
	}
	return value
}
