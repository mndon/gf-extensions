package httpx

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mndon/gf-extensions/errorx"
	"time"
)

const (
	nonceKey     = "nonce"
	timestampKey = "timestamp"
	signKey      = "sign"
)

var blacklist = gcache.New()

func getExpire() int64 {
	return g.Cfg().MustGet(context.TODO(), "safetyMiddleware.exp", 60*1000).Int64()
}

func getSalt() string {
	return g.Cfg().MustGet(context.TODO(), "safetyMiddleware.salt", "a947e5fd22f041cd82923fff4362169c").String()
}

func IsTestEnv() bool {
	return g.Cfg().MustGet(context.TODO(), "test", false).Bool()
}

func MiddlewareSafetyFilter(r *ghttp.Request) {
	encyRequest(r)
	r.Middleware.Next()
}

func encyRequest(r *ghttp.Request) {
	timestampStr := r.Header.Get(timestampKey)
	// 校验是时间字段
	if len(timestampStr) != 13 {
		SafetyFiltering(r.GetCtx(), "timestamp invalided")
	}
	// 校验签名时效性
	timestamp := gvar.New(timestampStr).Int64()
	now := time.Now().UnixMilli()
	if timestamp < (now - getExpire()) {
		SafetyFiltering(r.GetCtx(), "timestamp expired")
	}
	if timestamp > (now + 30000) { //冗余30秒
		g.Log().Warningf(r.GetCtx(), "timestamp invalided: timestamp: %+v, now: %+v", timestamp, now)
		SafetyFiltering(r.GetCtx(), "timestamp invalided")
	}
	// 校验nonce
	nonce := r.Header.Get(nonceKey)
	if len(nonce) != 32 {
		SafetyFiltering(r.GetCtx(), "nonce invalided")
	}
	// 校验nonce唯一性
	in, err := inBlacklist(r.GetCtx(), nonce)
	if err != nil {
		SafetyFiltering(r.GetCtx(), err.Error())
	}
	if in {
		SafetyFiltering(r.GetCtx(), "nonce invalided")
	}

	// 校验是签名字段
	sign := r.Header.Get(signKey)
	data := nonce + timestampStr + getSalt()
	calSign := md5Ency(data, "")
	if sign != calSign {
		msg := "Security Reason"
		if IsTestEnv() {
			msg = fmt.Sprintf("sign invalided, md5(%s)=%s", data, calSign)
		}
		SafetyFiltering(r.GetCtx(), msg)
	}

	err = setBlacklist(r.GetCtx(), nonce)
	if err != nil {
		SafetyFiltering(r.GetCtx(), err.Error())
	}
}

// SafetyFiltering 安全限制异常返回
func SafetyFiltering(ctx context.Context, msg string) {
	r := g.RequestFromCtx(ctx)
	r.SetError(errorx.NewErr(gcode.CodeSecurityReason.Code(), msg))
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeSecurityReason.Code(),
		Message: msg,
	})
	r.ExitAll()
}

// md5加密
func md5Ency(data string, salt string) string {
	h := md5.New()
	h.Write([]byte(data + salt))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

func setBlacklist(ctx context.Context, nonce string) error {
	err := blacklist.Set(ctx, nonce, true, time.Millisecond*time.Duration(getExpire()))

	if err != nil {
		return err
	}

	return nil
}

func inBlacklist(ctx context.Context, nonce string) (bool, error) {
	if in, err := blacklist.Contains(ctx, nonce); err != nil {
		return false, nil
	} else {
		return in, nil
	}
}
