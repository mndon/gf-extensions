package logx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

const CustomFieldsKey = "logCK"

type CustomFields struct {
	Uid        string        `json:",omitempty"` // 请求uid
	Type       string        `json:",omitempty"` // 日志类型 ACCESS / SLOW_ACCESS / 其他
	AccessTime time.Duration `json:",omitempty"` // 处理时长
	ResStatus  int           `json:",omitempty"` // 响应http状态码
	ReqMethod  string        `json:",omitempty"` // 请求方法
	ReqUri     string        `json:",omitempty"` // 请求uri
	ReqUrl     string        `json:",omitempty"` // 请求url
	ReqBody    string        `json:",omitempty"` // 请求body
	ReqIp      string        `json:",omitempty"` // 请求ip
	UA         string        `json:",omitempty"` // 请求agent
}

type Logger struct {
	*glog.Logger // Parent logger, if it is not empty, it means the logger is used in chaining function.
	Ctx          context.Context
	customFields *CustomFields
}

func New(ctx context.Context, name ...string) *Logger {
	v, ok := ctx.Value(CustomFieldsKey).(*CustomFields)
	if !ok {
		v = &CustomFields{}
		ctx = context.WithValue(ctx, CustomFieldsKey, v)
	}

	return &Logger{
		Logger:       g.Log(name...),
		Ctx:          ctx,
		customFields: v,
	}
}

// WithCustomFields
// @Description: 上下文设置log自定义字段
// @param ctx
// @param fields
// @return context.Context
func WithCustomFields(ctx context.Context, fields CustomFields) context.Context {
	v, ok := ctx.Value(CustomFieldsKey).(*CustomFields)
	if !ok {
		v = &CustomFields{}
		_ = gconv.Scan(gconv.String(fields), &v)
		return context.WithValue(ctx, CustomFieldsKey, v)
	} else {
		_ = gconv.Scan(gconv.String(fields), &v)
		return ctx
	}
}
