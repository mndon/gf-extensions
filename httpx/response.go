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
	Status int         `json:"status"    dc:"Error code"`
	Msg    string      `json:"msg" dc:"Error message"`
	Remark string      `json:"remark" dc:"client tip message"`
	Data   interface{} `json:"data"    dc:"Result data for certain request according API definition"`
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
