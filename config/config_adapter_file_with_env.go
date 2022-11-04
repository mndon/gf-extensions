package config

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/mndon/gf-extensions/config/utils"
)

type AdapterFileWithEnv struct {
	*gcfg.AdapterFile
}

// NewAdapterFileWithEnv
// @Description: 实例化AdapterFileWithEnv
// @return *AdapterFileWithEnv
func NewAdapterFileWithEnv() *AdapterFileWithEnv {
	adapter, err := gcfg.NewAdapterFile()
	if err != nil {
		panic("config error")
	}
	config := &AdapterFileWithEnv{}
	config.AdapterFile = adapter
	data, err := config.Data(context.TODO())
	if err != nil {
		panic("get config data error")
	}
	config.expandConfigWithEnv(data)
	return config
}

// init
// @Description: 初始化配置，解析配置中的环境变量
// @receiver c
// @param data
// @return map[string]interface{}
func (c *AdapterFileWithEnv) expandConfigWithEnv(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		if valueStr, ok := value.(string); ok {
			expandValue := utils.ExpandEnv(valueStr)
			data[key] = expandValue
		} else if valueStr, ok := value.(map[string]interface{}); ok {
			data[key] = c.expandConfigWithEnv(valueStr)
		}
	}
	return data
}
