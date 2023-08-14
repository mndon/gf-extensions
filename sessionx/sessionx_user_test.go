package sessionx

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

type User struct {
	Name string
}

type UserService struct {
}

func (u UserService) InternalGetUserByUid(ctx context.Context, uid string) (user any, err error) {
	return &User{Name: "name"}, nil
}

func TestSessionx(t *testing.T) {
	t.Run("get user", func(t *testing.T) {
		ctx := gctx.New()
		ctx.Value("gHttpRequestObject")

		InitSessionUtils(UserService{})
		user, err := GetUser[User](ctx)
		if err != nil {
			t.Error(err)
		}
		t.Log(user)
	})
}
