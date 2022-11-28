package http

import (
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/config"
	"time"
)

// jwt相关配置
const (
	JwtKey         = "jwt.key"
	JwtTimeout     = "jwt.timeout"
	JwtMaxRefresh  = "jwt.max_refresh"
	JwtIdentityKey = "jwt.identity_key"
)

var insJwt *jwt.GfJWTMiddleware

func init() {
	ctx := context.TODO()
	insJwt = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             config.GetValueFromConfigWithPanic(ctx, JwtKey, "HihfiasdnfdsnfsdnfiHNfikadnfknsd").Bytes(),
		Timeout:         time.Hour * config.GetValueFromConfigWithPanic(ctx, JwtTimeout, 24*30).Duration(),
		MaxRefresh:      time.Hour * config.GetValueFromConfigWithPanic(ctx, JwtMaxRefresh, 24*120).Duration(),
		IdentityKey:     config.GetValueFromConfigWithPanic(ctx, JwtIdentityKey, "uid").String(),
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
}

func JwtAuth() *jwt.GfJWTMiddleware {
	return insJwt
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[insJwt.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	responseCode := CodeAuthorizedErr
	r.Response.WriteJson(HandlerResponse{
		Status: responseCode.Code(),
		Msg:    "Invalid token",
		Remark: responseCode.Message(),
	})
	r.ExitAll()
}
