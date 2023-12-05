package jwtx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mndon/gf-extensions/errorx"
	"github.com/mndon/gf-extensions/sessionx"
	"github.com/mndon/gf-extensions/utilx"
	"net/http"
	"strings"
	"time"
)

const (
	JwtDefaultSecretKey = "HihfiasdnfdsnfsdnfiHNikaniki"
	TokenKey            = "JWT_TOKEN"
	PayloadKey          = "JWT_PAYLOAD"
	Exp                 = "exp"
	OrigIat             = "orig_iat"
)

type JwtOption struct {
	Timeout    time.Duration //jw过期时间
	MaxRefresh time.Duration //jwt最长刷新时间
	SecretKey  string        //jwt密钥

	LimitLogin  bool          // 限制所设备登录
	LimitCount  int           // 限制设备数
	CachePrefix string        //缓存key前缀
	Cache       *gcache.Cache //缓存

	RefreshTokenCustomerCheck func(ctx context.Context, uid string) error
}

type JwtAuth struct {
	IdentityKey      string // 用户标识字段名
	TokenLookup      string // token装载位置
	TokenHeadName    string // token头名称
	SigningAlgorithm string //加密方式

	BytesSecretKey []byte
	jwtOption      JwtOption
}

func NewJwt(opt JwtOption) *JwtAuth {
	if opt.SecretKey == "" {
		opt.SecretKey = JwtDefaultSecretKey
	}

	if opt.Timeout == 0 {
		opt.Timeout = time.Hour * 24
	}

	if opt.MaxRefresh == 0 {
		opt.MaxRefresh = time.Hour * 24 * 30
	}

	j := &JwtAuth{
		IdentityKey:      sessionx.CtxKeyUserUid,
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		SigningAlgorithm: "HS256",
		BytesSecretKey:   []byte(opt.SecretKey),
		jwtOption:        opt,
	}

	return j
}

// GenerateToken 构建token
func (j *JwtAuth) GenerateToken(ctx context.Context, userUid string) (token string, expire time.Time, err error) {
	token, t, err := j.buildJwt(g.Map{j.IdentityKey: userUid})
	if err != nil {
		return "", time.Time{}, err
	}
	if j.jwtOption.LimitLogin {
		// 将token塞入缓存
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
	// 校验过期
	r := g.RequestFromCtx(ctx)

	jwtToken, err := j.parseToken(r)
	if err != nil {
		// If we receive an error, and the error is anything other than a single
		// ValidationErrorExpired, we want to return the error.
		// If the error is just ValidationErrorExpired, we want to continue, as we can still
		// refresh the token if it's within the MaxRefresh time.
		// (see https://github.com/appleboy/gin-jwt/issues/176)
		validationErr, ok := err.(*jwt.ValidationError)
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return "", time.Time{}, err
		}
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	origIat := int64(claims[OrigIat].(float64))
	if origIat < time.Now().Add(-j.jwtOption.MaxRefresh).Unix() {
		return "", time.Time{}, errorx.BadRequestErr("token refresh is expired", errorx.CodeAuthorizedErr.Message())
	}

	// 校验用户有效性
	userUid := gconv.String(claims[j.IdentityKey])
	if userUid == "" {
		return "", time.Time{}, errorx.BadRequestErr("token invalid", errorx.CodeAuthorizedErr.Message())
	}
	if j.jwtOption.RefreshTokenCustomerCheck != nil {
		err := j.jwtOption.RefreshTokenCustomerCheck(ctx, userUid)
		if err != nil {
			return "", time.Time{}, err
		}
	}

	token, expire, err := j.buildJwt(g.Map{j.IdentityKey: userUid})
	if err != nil {
		return "", time.Time{}, err
	}

	// 添加至缓存
	if j.jwtOption.LimitLogin {
		err = j.addTokenToCache(ctx, userUid, token, expire)
		if err != nil {
			return "", time.Time{}, err
		}
	}
	return token, expire, nil
}

// MiddlewareJwtAuth 生成jwt鉴权中间件
func (j *JwtAuth) MiddlewareJwtAuth(r *ghttp.Request) {
	ctx := r.GetCtx()

	// 获取载荷
	claims, token, err := j.getClaimsFromJWT(ctx)
	if err != nil {
		j.unauthorized(ctx, errorx.NotAuthorizedErr("invalid token, code: 1"))
		return
	}
	// 获取过期时间
	if claims[Exp] == nil {
		j.unauthorized(ctx, errorx.NotAuthorizedErr("invalid token, code: 2"))
		return
	}
	if _, ok := claims[Exp].(float64); !ok {
		j.unauthorized(ctx, errorx.NotAuthorizedErr("invalid token, code: 3"))
		return
	}
	// 校验过期
	if int64(claims[Exp].(float64)) < time.Now().Unix() {
		j.unauthorized(ctx, errorx.NotAuthorizedErr("token expire"))
		return
	}

	// 校验用户id有效性
	userUid := gconv.String(claims[j.IdentityKey])
	if userUid == "" {
		j.unauthorized(ctx, errorx.NotAuthorizedErr("invalid token, code: 4"))
		return
	}

	r.SetParam(PayloadKey, claims)
	r.SetParam(j.IdentityKey, userUid)

	// 多设备限制登录
	if j.jwtOption.LimitLogin {
		userUid := sessionx.GetUserUid(ctx)
		cacheKey := j.buildCacheKey(userUid)
		v, err := j.jwtOption.Cache.Get(ctx, cacheKey)
		if err != nil {
			j.unauthorized(ctx, errorx.NewErrWithCode(errorx.CodeInternalErr, "Login limit cache error, code: 1"))
			return
		}
		tokens := v.Strings()
		// tokens在限制数量内
		if len(tokens) < j.jwtOption.LimitCount {
			// token不在缓存中，则新增
			if !utilx.StringsIn(token, tokens) {
				tokenExpireTime := time.Unix(int64(claims[Exp].(float64)), 0)
				err = j.addTokenToCache(ctx, userUid, token, tokenExpireTime)
				if err != nil {
					j.unauthorized(ctx, errorx.NewErrWithCode(errorx.CodeInternalErr, "Login limit cache error, code: 2"))
					return
				}
			}
		} else {
			// tokens到达限制数量内，且token不存在，则推出登陆
			if !utilx.StringsIn(token, tokens) {
				j.unauthorized(ctx, errorx.NewErrWithCode(errorx.CodeLoginLimitErr, "login of multiple devices"))
				return
			}
		}
	}

	r.Middleware.Next()
}

// unauthorized
// @Description: 鉴权失败回调
// @receiver j
// @param ctx
// @param code
// @param message
func (j *JwtAuth) unauthorized(ctx context.Context, err error) {
	r := g.RequestFromCtx(ctx)
	r.SetError(err)
	r.Response.Status = http.StatusOK
}

// TokenGenerator 构建jwt.
func (j *JwtAuth) buildJwt(payload map[string]any) (token string, expire time.Time, err error) {
	jwtToken := jwt.New(jwt.GetSigningMethod(j.SigningAlgorithm))
	claims := jwtToken.Claims.(jwt.MapClaims)
	if payload != nil {
		for key, value := range payload {
			claims[key] = value
		}
	}

	now := time.Now()
	expire = now.Add(j.jwtOption.Timeout)
	claims[Exp] = expire.Unix()
	claims[OrigIat] = now.Unix()
	token, err = j.signedString(jwtToken)
	if err != nil {
		return "", time.Time{}, err
	}
	return token, expire, nil
}

// getClaimsFromJWT 从jwt获取载荷
func (j *JwtAuth) getClaimsFromJWT(ctx context.Context) (claims jwt.MapClaims, token string, err error) {
	r := g.RequestFromCtx(ctx)

	jwtToken, err := j.parseToken(r)
	if err != nil {
		return nil, "", err
	}

	return jwtToken.Claims.(jwt.MapClaims), jwtToken.Raw, nil
}

// 解析token
func (j *JwtAuth) parseToken(r *ghttp.Request) (*jwt.Token, error) {
	var token string
	var err error

	methods := strings.Split(j.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = j.jwtFromHeader(r, v)
		case "query":
			token, err = j.jwtFromQuery(r, v)
		case "cookie":
			token, err = j.jwtFromCookie(r, v)
		case "param":
			token, err = j.jwtFromParam(r, v)
		}
	}

	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(j.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		r.SetParam(TokenKey, token)
		return j.BytesSecretKey, nil
	})
}

func (j *JwtAuth) jwtFromHeader(r *ghttp.Request, key string) (string, error) {
	authHeader := r.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == j.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func (j *JwtAuth) jwtFromQuery(r *ghttp.Request, key string) (string, error) {
	token := r.Get(key).String()

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func (j *JwtAuth) jwtFromCookie(r *ghttp.Request, key string) (string, error) {
	cookie := r.Cookie.Get(key).String()

	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

func (j *JwtAuth) jwtFromParam(r *ghttp.Request, key string) (string, error) {
	token := r.Get(key).String()

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}

func (j *JwtAuth) signedString(token *jwt.Token) (string, error) {
	return token.SignedString(j.BytesSecretKey)
}

func (j *JwtAuth) buildCacheKey(key string) string {
	return "JWT." + j.jwtOption.CachePrefix + "." + key
}

func (j *JwtAuth) addTokenToCache(ctx context.Context, userUid string, token string, expireTime time.Time) error {
	// 限制多设备登录
	cacheKey := j.buildCacheKey(userUid)
	v, err := j.jwtOption.Cache.Get(ctx, cacheKey)
	if err != nil {
		return err
	}
	tokens := v.Strings()
	tokens = append([]string{token}, tokens...)
	if len(tokens) > j.jwtOption.LimitCount {
		tokens = tokens[:j.jwtOption.LimitCount]
	}

	err = j.jwtOption.Cache.Set(context.TODO(), cacheKey, tokens, expireTime.Sub(time.Now()))
	if err != nil {
		return err
	}
	return nil
}
