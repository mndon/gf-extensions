package envx

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
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
		env := gcmd.GetOpt("MY_ENV", "").String()
		if env == "" {
			env = genv.Get("MY_ENV", EnvLocal).String()
		}
		env = strings.ToLower(env)
		if env != EnvLocal && env != EnvTest && env != EnvStaging && env != EnvProd {
			env = EnvLocal
		}
		Env = env
		fmt.Println("MY_ENV:", Env)
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
