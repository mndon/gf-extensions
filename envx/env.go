package envx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"sync"
)

var Env string
var initEnv sync.Once

const (
	EnvLocal   = "local"   // 本地
	EnvTest    = "test"    // 测试
	EnvStaging = "staging" // 预生产
	EnvProd    = "prod"    // 生产
)

func GetEnv() string {
	initEnv.Do(func() {
		env := g.Cfg().MustGet(context.TODO(), "MY_ENV", EnvLocal).String()
		env = strings.ToLower(env)
		if env != EnvLocal && env != EnvTest && env != EnvStaging && env != EnvProd {
			env = EnvLocal
		}
		Env = env
	})
	return Env
}

func IsLocal() bool {
	return GetEnv() == EnvLocal
}

func IsTest() bool {
	return GetEnv() == EnvTest
}

func IsStaging() bool {
	return GetEnv() == EnvStaging
}

func IsProd() bool {
	return GetEnv() == EnvProd
}
