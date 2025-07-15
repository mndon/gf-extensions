/*
* @desc:xxxx功能描述
* @company:
* @Author:
* @Date:   2022/11/3 9:55
 */

package personal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/libUtils"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterPersonal(New())
}

type sPersonal struct {
}

func New() *sPersonal {
	return &sPersonal{}
}

func (s *sPersonal) GetPersonalInfo(ctx context.Context, req *api.PersonalInfoReq) (res *api.PersonalInfoRes, err error) {
	res = new(api.PersonalInfoRes)
	userId := service.Context().GetUserId(ctx)
	res.User, err = service.AdminUser().GetUserInfoById(ctx, userId)
	allRoles, err := service.AdminRole().GetRoleList(ctx)
	roles, err := service.AdminUser().GetAdminRole(ctx, userId, allRoles)
	name := make([]string, len(roles))
	roleIds := make([]uint, len(roles))
	for k, v := range roles {
		name[k] = v.Name
		roleIds[k] = v.Id
	}
	res.Roles = name
	if err != nil {
		return
	}
	return
}

func (s *sPersonal) EditPersonal(ctx context.Context, req *api.PersonalEditReq) (user *model.LoginUserRes, err error) {
	userId := service.Context().GetUserId(ctx)
	err = service.AdminUser().UserNameOrMobileExists(ctx, "", req.Mobile, int64(userId))
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.AdminUser.Ctx(ctx).TX(tx).WherePri(userId).Update(do.AdminUser{
				Mobile:       req.Mobile,
				UserNickname: req.Nickname,
				Remark:       req.Remark,
				Sex:          req.Sex,
				UserEmail:    req.UserEmail,
				Describe:     req.Describe,
				Avatar:       req.Avatar,
			})
			liberr.ErrIsNil(ctx, err, "修改用户信息失败")
			user, err = service.AdminUser().GetUserById(ctx, userId)
			liberr.ErrIsNil(ctx, err)
		})
		return err
	})
	return
}

func (s *sPersonal) ResetPwdPersonal(ctx context.Context, req *api.PersonalResetPwdReq) (res *api.PersonalResetPwdRes, err error) {
	userId := service.Context().GetUserId(ctx)
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminUser.Ctx(ctx).WherePri(userId).Update(g.Map{
			dao.AdminUser.Columns().UserSalt:     salt,
			dao.AdminUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "重置用户密码失败")
	})
	return
}
