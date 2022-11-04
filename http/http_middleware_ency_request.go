package http

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sort"
	"strconv"
	"time"
)

// MiddlewareEncyVerif 验证请求签名
func MiddlewareEncyVerif(r *ghttp.Request) {
	timestampInputStr := r.Header.Get("timestamp")
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
	signInput := r.Header.Get("sign")
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
		data = data + k + reqMap[k]
	}
	// 加时间戳字符串
	data = data + timestampInputStr
	// 加盐
	salt := "a947e5fd22f041cd82923fff4362169c"
	sign := md5Ency(data, salt)
	if sign != signInput {
		g.Log().Warning(r.GetCtx(), fmt.Sprintf("sign invalided, data: %s, right sign: %s", data, sign))
		SafeFiltering(r.GetCtx())
	}
	r.Middleware.Next()
}

func md5Ency(data string, salt string) string {
	h := md5.New()
	h.Write([]byte(data + salt))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}
