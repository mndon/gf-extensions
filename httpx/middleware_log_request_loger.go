package httpx

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var accessLogger *glog.Logger

func init() {
	config := g.Log().GetConfig()
	config.StStatus = 0
	accessLogger = glog.New()
	accessLogger.SetConfig(config)
}
