package http

import (
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"time"
)

const (
	JwtDefaultSecretKey = "HihfiasdnfdsnfsdnfiHNikaniki"
	JwtIdentityKey      = "uid" // 固定载荷 IdentityKey 为uid
)

type Jwt struct {
	IdentityKey   string
	TokenLookup   string
	TokenHeadName string
	secretKey     string
	jm            *jwt.GfJWTMiddleware
}

func NewJwt(timeout int, maxRefresh int, jwtKey ...string) *Jwt {
	timeoutDuration := time.Hour * time.Duration(timeout)
	maxRefreshDuration := time.Hour * time.Duration(maxRefresh)
	return NewJwtWithTimeDuration(timeoutDuration, maxRefreshDuration, jwtKey...)
}

func NewJwtWithTimeDuration(timeout time.Duration, maxRefresh time.Duration, jwtKey ...string) *Jwt {
	key := JwtDefaultSecretKey
	if len(jwtKey) == 1 {
		key = jwtKey[0]
	}

	j := &Jwt{
		IdentityKey:   JwtIdentityKey,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		secretKey:     key,
	}

	j.jm = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "jwt",
		Key:             []byte(j.secretKey),
		Timeout:         timeout,
		MaxRefresh:      maxRefresh,
		IdentityKey:     j.IdentityKey,
		TokenLookup:     j.TokenLookup,
		TokenHeadName:   j.TokenHeadName,
		TimeFunc:        time.Now,
		Unauthorized:    j.unauthorized,
		PayloadFunc:     j.payloadFunc,
		IdentityHandler: j.identityHandler,
	})

	return j
}

// TokenGenerator
// @Description: token生成
// @receiver j
// @param userUid
// @return string
// @return time.Time
// @return error
func (j *Jwt) TokenGenerator(userUid string) (string, time.Time, error) {
	return j.jm.TokenGenerator(g.Map{j.IdentityKey: userUid})
}

// RefreshToken
// @Description: token刷新
// @receiver j
// @param ctx
// @return string
// @return time.Time
// @return error
func (j Jwt) RefreshToken(ctx context.Context) (string, time.Time, error) {
	return j.jm.RefreshToken(ctx)
}

// GetIdentity
// @Description: 获取jwt载荷中Identity字段的值
// @receiver j
// @param ctx
// @return interface{}
func (j *Jwt) GetIdentity(ctx context.Context) interface{} {
	return j.jm.GetIdentity(ctx)
}

// BuildMiddlewareJwtAuth
// @Description: 生成jwt鉴权中间件
// @receiver j
// @param r
func (j *Jwt) BuildMiddlewareJwtAuth(r *ghttp.Request) {
	j.jm.MiddlewareFunc()(r)
	r.Middleware.Next()
}

// payloadFunc
// @Description: 载荷方法
// @param data
// @return jwt.MapClaims
func (j *Jwt) payloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// identityHandler
// @Description: 身份验证回调
// @param ctx
// @return interface{}
func (j *Jwt) identityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	i := claims[j.jm.IdentityKey]
	if i == nil || i == "" || i == 0 {
		return nil
	}
	return i
}

// unauthorized
// @Description: 鉴权失败回调
// @receiver j
// @param ctx
// @param code
// @param message
func (j *Jwt) unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.SetError(NotAuthorizedErr("Invalid token"))
	r.Response.Status = http.StatusOK
}

// GetIdentityFromJwtFromCtx
// @Description: 从上下文中获取jwt携带得uid
// @param ctx
// @return string
func GetIdentityFromJwtFromCtx(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	return r.Get(JwtIdentityKey).String()
}

// SetIdentityInCtx
// @Description: 设置上下文中的uid
// @param ctx
// @return string
func SetIdentityInCtx(ctx context.Context, uid string) {
	r := g.RequestFromCtx(ctx)
	r.SetParam(JwtIdentityKey, uid)
}
