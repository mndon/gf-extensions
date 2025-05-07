package httpx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	ctxKeyForRemark = "__internal_http_response_remark"
)

// HandlerResponse
// @Description: 响应结构体
type HandlerResponse struct {
	Status  int         `json:"status"  `
	Msg     string      `json:"msg"`
	Remark  string      `json:"remark"`
	TraceId string      `json:"trace_id"`
	Data    interface{} `json:"data"`
}

// GetResponseRemark
// @Description: 获取响应体体提示内容
// @param ctx
// @return string
func GetResponseRemark(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).GetCtxVar(ctxKeyForRemark).String()
}

// SetResponseRemark
// @Description: 设置响应体提示内容
// @param ctx
// @param remark
func SetResponseRemark(ctx context.Context, remark string) {
	ghttp.RequestFromCtx(ctx).SetCtxVar(ctxKeyForRemark, remark)
}
