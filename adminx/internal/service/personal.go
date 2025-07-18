// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/model"
)

type (
	IPersonal interface {
		GetPersonalInfo(ctx context.Context, req *api.PersonalInfoReq) (res *api.PersonalInfoRes, err error)
		EditPersonal(ctx context.Context, req *api.PersonalEditReq) (user *model.LoginUserRes, err error)
		ResetPwdPersonal(ctx context.Context, req *api.PersonalResetPwdReq) (res *api.PersonalResetPwdRes, err error)
	}
)

var (
	localPersonal IPersonal
)

func Personal() IPersonal {
	if localPersonal == nil {
		panic("implement not found for interface IPersonal, forgot register?")
	}
	return localPersonal
}

func RegisterPersonal(i IPersonal) {
	localPersonal = i
}
