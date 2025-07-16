package adminx

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

// GetUser
// @Description: 获取后台用户信息
// @param ctx
// @return *model.LoginUserRes
func GetUser(ctx context.Context) *model.LoginUserRes {
	return service.Context().Get(ctx).User.LoginUserRes
}
