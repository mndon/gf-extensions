package http

import (
	"context"
	gjwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const (
	JwtDefaultSecretKey = "HihfiasdnfdsnfsdnfiHNikaniki"
	JwtIdentityKey      = "uid" // 固定载荷 IdentityKey 为uid
)

type JwtAuthLoginLimitOption struct {
	LimitLogin  bool // 限制所设备登录
	LimitCount  int  // 限制设备数
	CachePrefix string
	Cache       *gcache.Cache //缓存
}

type JwtAuth struct {
	IdentityKey   string
	TokenLookup   string
	TokenHeadName string
	secretKey     string
	mw            *gjwt.GfJWTMiddleware

	loginLimitOption *JwtAuthLoginLimitOption
}

func NewJwtAuth(timeout int, maxRefresh int, loginLimitOption *JwtAuthLoginLimitOption, jwtKey ...string) *JwtAuth {
	timeoutDuration := time.Hour * time.Duration(timeout)
	maxRefreshDuration := time.Hour * time.Duration(maxRefresh)
	return NewJwtWithTimeDuration(timeoutDuration, maxRefreshDuration, loginLimitOption, jwtKey...)
}

func NewJwtWithTimeDuration(timeout time.Duration, maxRefresh time.Duration, loginLimitOption *JwtAuthLoginLimitOption, jwtKey ...string) *JwtAuth {
	key := JwtDefaultSecretKey
	if len(jwtKey) == 1 {
		key = jwtKey[0]
	}

	j := &JwtAuth{
		IdentityKey:      JwtIdentityKey,
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		secretKey:        key,
		loginLimitOption: loginLimitOption,
	}

	j.mw = gjwt.New(&gjwt.GfJWTMiddleware{
		Realm:           "jwt",
		Key:             []byte(j.secretKey),
		Timeout:         timeout,
		MaxRefresh:      maxRefresh,
		IdentityKey:     j.IdentityKey,
		TokenLookup:     j.TokenLookup,
		TokenHeadName:   j.TokenHeadName,
		TimeFunc:        time.Now,
		Authorizator:    j.authorized,
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
func (j *JwtAuth) TokenGenerator(ctx context.Context, userUid string) (string, time.Time, error) {
	token, t, err := j.buildJwt(g.Map{j.IdentityKey: userUid})
	if err != nil {
		return "", time.Time{}, err
	}
	if j.loginLimitOption.LimitLogin {
		err = j.addTokenToCache(ctx, userUid, token, t)
		if err != nil {
			return "", time.Time{}, err
		}
	}
	return token, t, nil
}

// RefreshToken
// @Description: token刷新
// @receiver j
// @param ctx
// @return string
// @return time.Time
// @return error
func (j *JwtAuth) RefreshToken(ctx context.Context) (string, time.Time, error) {
	// 刷新token
	claims, _, err := j.mw.CheckIfTokenExpire(ctx)
	if err != nil {
		return "", time.Now(), err
	}
	authorized := j.authorized(claims[j.IdentityKey], ctx)
	if !authorized {
		return "", time.Time{}, NotAuthorizedErr("token invalid")
	}

	userUid := gconv.String(claims[j.IdentityKey])
	token, expire, err := j.buildJwt(g.Map{j.IdentityKey: userUid})
	if err != nil {
		return "", time.Time{}, err
	}

	// 添加至缓存
	if j.loginLimitOption.LimitLogin {
		err = j.addTokenToCache(ctx, userUid, token, expire)
		if err != nil {
			return "", time.Time{}, err
		}
	}
	return token, expire, nil
}

// TokenGenerator method that clients can use to get a jwt token.
func (j *JwtAuth) buildJwt(data interface{}) (string, time.Time, error) {
	token := jwt.New(jwt.GetSigningMethod(j.mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if j.mw.PayloadFunc != nil {
		for key, value := range j.mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := time.Now().UTC().Add(j.mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()
	tokenString, err := j.signedString(token)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expire, nil
}

func (j *JwtAuth) signedString(token *jwt.Token) (string, error) {
	return token.SignedString(j.mw.Key)
}

func (j *JwtAuth) buildCacheKey(key string) string {
	return "JWT." + j.loginLimitOption.CachePrefix + "." + key
}

func (j *JwtAuth) addTokenToCache(ctx context.Context, userUid string, token string, expireTime time.Time) error {
	// 限制多设备登录
	cacheKey := j.buildCacheKey(userUid)
	l, err := j.loginLimitOption.Cache.Get(ctx, cacheKey)
	if err != nil {
		return err
	}
	tokens := l.Strings()
	tokens = append([]string{token}, tokens...)
	if len(tokens) > j.loginLimitOption.LimitCount {
		tokens = tokens[:j.loginLimitOption.LimitCount]
	}

	err = j.loginLimitOption.Cache.Set(context.TODO(), cacheKey, tokens, expireTime.Sub(time.Now()))
	if err != nil {
		return err
	}
	return nil
}

// GetIdentity
// @Description: 获取jwt载荷中Identity字段的值
// @receiver j
// @param ctx
// @return interface{}
func (j *JwtAuth) GetIdentity(ctx context.Context) interface{} {
	return j.mw.GetIdentity(ctx)
}

// BuildMiddlewareJwtAuth
// @Description: 生成jwt鉴权中间件
// @receiver j
// @param r
func (j *JwtAuth) BuildMiddlewareJwtAuth(r *ghttp.Request) {
	// jwt校验
	j.mw.MiddlewareFunc()(r)

	// 多设备限制登录
	if j.loginLimitOption.LimitLogin {
		userUid := gconv.String(j.GetIdentity(r.GetCtx()))
		cacheKey := j.buildCacheKey(userUid)
		l, err := j.loginLimitOption.Cache.Get(r.GetCtx(), cacheKey)
		if err != nil {
			j.unauthorized(r.GetCtx(), http.StatusInternalServerError, "Login limit Cache exception")
			return
		}
		tokens := l.Strings()
		c, token, err := j.mw.GetClaimsFromJWT(r.GetCtx())
		if len(tokens) < j.loginLimitOption.LimitCount {
			if err != nil {
				j.unauthorized(r.GetCtx(), http.StatusInternalServerError, "Login limit parse token error")
				return
			}
			tokenExpireTime := time.Unix(int64(c["exp"].(float64)), 0)
			err = j.addTokenToCache(r.GetCtx(), userUid, token, tokenExpireTime)
			if err != nil {
				j.unauthorized(r.GetCtx(), http.StatusInternalServerError, "Login limit parse token error")
				return
			}
		} else {
			tokenExist := false
			for _, t := range tokens {
				if t == token {
					tokenExist = true
				}
			}
			// token数量超过限制且token不存在
			if !tokenExist {
				r.SetError(NewErrWithCode(LoginLimitErr, "login of multiple devices"))
				return
			}
		}
	}

	r.Middleware.Next()
}

// payloadFunc
// @Description: 载荷方法
// @param data
// @return jwt.MapClaims
func (j *JwtAuth) payloadFunc(data interface{}) gjwt.MapClaims {
	claims := gjwt.MapClaims{}
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
func (j *JwtAuth) identityHandler(ctx context.Context) interface{} {
	claims := gjwt.ExtractClaims(ctx)
	i := claims[j.mw.IdentityKey]
	if i == nil || i == "" || i == 0 {
		return nil
	}
	return i
}

func (j *JwtAuth) authorized(data interface{}, ctx context.Context) bool {
	iStr := gvar.New(data).String()
	if len(iStr) != 32 {
		return false
	}
	return true
}

// unauthorized
// @Description: 鉴权失败回调
// @receiver j
// @param ctx
// @param code
// @param message
func (j *JwtAuth) unauthorized(ctx context.Context, code int, message string) {
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
