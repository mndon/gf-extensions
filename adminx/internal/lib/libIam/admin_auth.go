package libIam

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var authURL string

func init() {
	authURL = g.Cfg().MustGet(gctx.New(), "iamUrl", "http://116.62.63.179:8090/user-info/").String()
}

type ValidUserPasswordReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type ValidUserPasswordRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

// ValidUserPassword
// @Description: 判断用户是否有效
// @param ctx
// @param username
// @param password
// @return err
func ValidUserPassword(ctx context.Context, username, password string) (err error) {
	res := ValidUserPasswordRes{}
	r := g.Client().PostVar(ctx, authURL, ValidUserPasswordReq{
		Username: username,
		Password: password,
	})
	err = r.Scan(&res)
	if err != nil {
		return err
	}
	if res.Status != 2000 {
		return gerror.New("auth error")
	}
	return nil
}
