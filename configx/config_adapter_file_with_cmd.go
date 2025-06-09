package configx

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/mndon/gf-extensions/envx"
)

type AdapterFileWithCmd struct {
	adapterFile *gcfg.AdapterFile
}

// NewAdapterFileWithCmd
// @Description: 实例化AdapterFileWithEnv
// @return *AdapterFileWithEnv
func NewAdapterFileWithCmd() *AdapterFileWithCmd {
	adapter, err := gcfg.NewAdapterFile()
	if err != nil {
		panic("config error")
	}
	env := envx.GetEnv()
	switch env {
	case envx.EnvProd:
		adapter.SetFileName("config_prod.yaml")
	case envx.EnvStaging:
		adapter.SetFileName("config_staging.yaml")
	case envx.EnvTest:
		adapter.SetFileName("config_test.yaml")
	case envx.EnvLocal:
		adapter.SetFileName("config.yaml")
	default:
		adapter.SetFileName("config.yaml")
	}

	config := &AdapterFileWithCmd{}
	config.adapterFile = adapter
	data, err := config.Data(context.TODO())
	if err != nil {
		panic("get config data error")
	}
	config.expandConfigWithCmd(data)
	return config
}

// Available checks and returns whether configuration of given `file` is available.
func (a *AdapterFileWithCmd) Available(ctx context.Context, resource ...string) bool {
	return a.adapterFile.Available(ctx, resource...)
}

// Data retrieves and returns all configuration data as map type.
func (a *AdapterFileWithCmd) Data(ctx context.Context) (data map[string]interface{}, err error) {
	return a.adapterFile.Data(ctx)
}

func (a *AdapterFileWithCmd) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	return a.adapterFile.Get(ctx, pattern)
}

func (a *AdapterFileWithCmd) expandConfigWithCmd(data map[string]interface{}) map[string]interface{} {
	cmdOptMap := gcmd.GetOptAll()
	for k, v := range cmdOptMap {
		err := a.adapterFile.Set(k, v)
		if err != nil {
			panic(err)
		}
	}
	return data
}
