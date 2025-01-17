package logx

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

const CustomFieldsKey = "logCK"

type customFields struct {
	Type       string        `json:",omitempty"` // 日志类型 ACCESS / SLOW_ACCESS / 其他
	AccessTime time.Duration `json:",omitempty"` // 处理时长
	ReqStatus  int           `json:",omitempty"` // 响应http状态码
	ReqMethod  string        `json:",omitempty"` // 请求方法
	ReqUri     string        `json:",omitempty"` // 请求uri
	ReqUrl     string        `json:",omitempty"` // 请求url
	ReqBody    string        `json:",omitempty"` // 请求body
	ReqIp      string        `json:",omitempty"` // 请求ip
	UA         string        `json:",omitempty"` // 请求agent
	Uid        string        `json:",omitempty"` // 请求uid
}

type Logger struct {
	*glog.Logger // Parent logger, if it is not empty, it means the logger is used in chaining function.
	customFields customFields
}

func New(name ...string) *Logger {
	return &Logger{
		Logger:       g.Log(name...),
		customFields: customFields{},
	}
}
