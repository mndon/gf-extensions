package sessionx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/errorx"
)

const (
	CtxKeyUserUid = "__internal_session_user_uid" // 固定载荷
	CtxKeyUser    = "__internal_session_user"
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
	return r.Get(CtxKeyUserUid).String()
}

// SetUserUid 设置用户uid
func SetUserUid(ctx context.Context, uid string) {
	r := g.RequestFromCtx(ctx)
	r.SetParam(CtxKeyUserUid, uid)
}

// GetUserWithOpt 获取用户
// todo
func GetUserWithOpt[T any](ctx context.Context, opt Opt) (user *T, err error) {
	r := g.RequestFromCtx(ctx)
	// 获取缓存
	if opt.Refresh == false {
		var u *T
		err = r.Get(CtxKeyUser).Scan(&u)
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
		return nil, errorx.InternalErr("get user uid error")
	}
	newU, err := sessionUtils.userService.InternalGetUserByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	switch v := newU.(type) {
	case *T:
		// 刷新缓存
		r.SetParam(CtxKeyUser, newU)
		return v, nil
	default:
		return nil, errorx.InternalErr("get user error")
	}
}

// GetUser 获取用户
// todo
func GetUser[T any](ctx context.Context) (user *T, err error) {
	return GetUserWithOpt[T](ctx, Opt{Refresh: false})
}
