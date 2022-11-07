package http

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sort"
	"strconv"
	"time"
)

const (
	timestampKey = "timestamp"
	signKey      = "sign"
	defaultSalt  = "a947e5fd22f041cd82923fff4362169c"
)

// EncyRequestMiddleware 验证请求签名中间件
func EncyRequestMiddleware(r *ghttp.Request) {
	encyRequest(r)
	r.Middleware.Next()
}

func encyRequest(r *ghttp.Request) {
	timestampInputStr := r.Header.Get(timestampKey)
	// 校验是时间字段
	if timestampInputStr == "" {
		g.Log().Info(r.GetCtx(), r.RemoteAddr+"no timestampInputStr")
		SafeFiltering(r.GetCtx())
	}
	// 校验签名时效性
	timestampInput, err := strconv.Atoi(timestampInputStr)
	if err != nil {
		g.Log().Info(r.GetCtx(), r.RemoteAddr+" timestampInput change to int error:"+timestampInputStr)
		SafeFiltering(r.GetCtx())
	} else if int64(timestampInput)+int64(15*24*60*1000) < time.Now().UnixMilli() {
		g.Log().Info(r.GetCtx(), r.RemoteAddr+" timestampInput expired: "+timestampInputStr)
		SafeFiltering(r.GetCtx())
	}
	// 校验是签名字段
	signInput := r.Header.Get(signKey)
	if signInput == "" {
		g.Log().Info(r.GetCtx(), r.RemoteAddr+"no sign")
		SafeFiltering(r.GetCtx())
	}

	// 获取参数
	data := ""
	reqMap := r.GetRequestMapStrStr()
	routerMap := r.GetRouterMap()
	var reqKeys []string
	// 过滤路径参数
	for k := range reqMap {
		_, ok := routerMap[k]
		if !ok {
			reqKeys = append(reqKeys, k)
		}
	}
	// 按key升序排序
	sort.Strings(reqKeys)
	// 拼接key value
	for _, k := range reqKeys {
		if reqMap[k] != "" {
			data = data + k + reqMap[k]
		}
	}
	// 加时间戳字符串
	data = data + timestampInputStr
	// 加盐
	sign := Utils().md5Ency(data, defaultSalt)
	if sign != signInput {
		g.Log().Warning(r.GetCtx(), fmt.Sprintf("sign invalided, data: %s, right sign: %s", data, sign))
		SafeFiltering(r.GetCtx())
	}
}

// SafeFiltering 安全限制异常返回
func SafeFiltering(ctx context.Context) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeSecurityReason.Code(),
		Message: gcode.CodeSecurityReason.Message(),
	})
	r.ExitAll()
}
