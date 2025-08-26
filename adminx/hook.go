package adminx

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

// OperationLog
// @Description: 记录操作日志
// @param r
func OperationLog(r *ghttp.Request) {
	service.OperateLog().OperationLog(r)
}
