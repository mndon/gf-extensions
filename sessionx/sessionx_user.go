package sessionx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/httpx"
)

const (
	UserUid    = "uid" // 固定载荷
	ctxUserKey = "__user"
)

type IUserService interface {
	InternalGetUserByUid(ctx context.Context, uid string) (user any, err error)
}

type SessionUtils struct {
	userService IUserService
}

type Opt struct {
	Refresh bool
}

var sessionUtils *SessionUtils

func InitSessionUtils(service IUserService) {
	sessionUtils = &SessionUtils{userService: service}
}

// GetUserUid 获取用户uid
func GetUserUid(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	return r.Get(UserUid).String()
}

// SetUserUid 设置用户uid
func SetUserUid(ctx context.Context, uid string) {
	r := g.RequestFromCtx(ctx)
	r.SetParam(UserUid, uid)
}

// GetUserWithOpt 获取用户
func GetUserWithOpt[T any](ctx context.Context, opt Opt) (user *T, err error) {
	r := g.RequestFromCtx(ctx)
	// 获取缓存
	if opt.Refresh == false {
		var u *T
		err = r.Get(ctxUserKey).Scan(&u)
		if err != nil {
			return nil, err
		}
		if u != nil {
			return u, nil
		}
	}
	// 获取用户
	uid := GetUserUid(ctx)
	if uid == "" {
		return nil, httpx.InternalErr("get user uid error")
	}
	newU, err := sessionUtils.userService.InternalGetUserByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	switch v := newU.(type) {
	case *T:
		// 刷新缓存
		r.SetParam(ctxUserKey, newU)
		return v, nil
	default:
		return nil, httpx.InternalErr("get user error")
	}
}

// GetUser 获取用户
func GetUser[T any](ctx context.Context) (user *T, err error) {
	return GetUserWithOpt[T](ctx, Opt{Refresh: false})
}
