/*
* @desc:用户在线状态处理
* @company:
* @Author:
* @Date:   2023/1/10 14:50
 */

package adminUserOnline

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/mssola/user_agent"
)

func init() {
	service.RegisterAdminUserOnline(New())
}

func New() *sAdminUserOnline {
	return &sAdminUserOnline{
		Pool: grpool.New(100),
	}
}

type sAdminUserOnline struct {
	Pool *grpool.Pool
}

func (s *sAdminUserOnline) Invoke(ctx context.Context, params *model.AdminUserOnlineParams) {
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入数据
		s.SaveOnline(ctx, params)
	})
}

// SaveOnline 保存用户在线状态
func (s *sAdminUserOnline) SaveOnline(ctx context.Context, params *model.AdminUserOnlineParams) {
	err := g.Try(ctx, func(ctx context.Context) {
		ua := user_agent.New(params.UserAgent)
		browser, _ := ua.Browser()
		os := ua.OS()
		var (
			info *entity.AdminUserOnline
			data = &do.AdminUserOnline{
				Uuid:       params.Uuid,
				Token:      params.Token,
				CreateTime: gtime.Now(),
				UserName:   params.Username,
				Ip:         params.Ip,
				Explorer:   browser,
				Os:         os,
			}
		)

		//查询是否已存在当前用户
		err := dao.AdminUserOnline.Ctx(ctx).Fields(dao.AdminUserOnline.Columns().Id).
			Where(dao.AdminUserOnline.Columns().Token, data.Token).
			Scan(&info)
		liberr.ErrIsNil(ctx, err)
		//若已存在则更新
		if info != nil {
			_, err = dao.AdminUserOnline.Ctx(ctx).
				Where(dao.AdminUserOnline.Columns().Id, info.Id).
				FieldsEx(dao.AdminUserOnline.Columns().Id).Update(data)
			liberr.ErrIsNil(ctx, err)
		} else { //否则新增
			_, err = dao.AdminUserOnline.Ctx(ctx).
				FieldsEx(dao.AdminUserOnline.Columns().Id).Insert(data)
			liberr.ErrIsNil(ctx, err)
		}
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// CheckUserOnline 检查在线用户
func (s *sAdminUserOnline) CheckUserOnline(ctx context.Context) {
	param := &api.AdminUserOnlineSearchReq{
		PageReq: api.PageReq{
			PageReq: model.PageReq{
				PageNum:  1,
				PageSize: 50,
			},
		},
	}
	var total int
	for {
		var (
			res *api.AdminUserOnlineSearchRes
			err error
		)
		res, err = s.GetOnlineListPage(ctx, param, true)
		if err != nil {
			g.Log().Error(ctx, err)
			break
		}
		if res.List == nil {
			break
		}
		for _, v := range res.List {
			if b := s.UserIsOnline(ctx, v.Token); !b {
				s.DeleteOnlineByToken(ctx, v.Token)
			}
		}
		if param.PageNum*param.PageSize >= total {
			break
		}
		param.PageNum++
	}
}

// GetOnlineListPage 搜素在线用户列表
func (s *sAdminUserOnline) GetOnlineListPage(ctx context.Context, req *api.AdminUserOnlineSearchReq, hasToken ...bool) (res *api.AdminUserOnlineSearchRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	model := dao.AdminUserOnline.Ctx(ctx)
	if req.Ip != "" {
		model = model.Where("ip like ?", "%"+req.Ip+"%")
	}
	if req.Username != "" {
		model = model.Where("user_name like ?", "%"+req.Username+"%")
	}
	res = new(api.AdminUserOnlineSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if len(hasToken) == 0 || !hasToken[0] {
			model = model.FieldsEx("token")
		}
		err = model.Page(req.PageNum, req.PageSize).Order("create_time DESC").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sAdminUserOnline) UserIsOnline(ctx context.Context, token string) bool {
	err := g.Try(ctx, func(ctx context.Context) {
		_, _, err := service.GfToken().GetTokenData(ctx, token)
		liberr.ErrIsNil(ctx, err)
	})
	return err == nil
}

func (s *sAdminUserOnline) DeleteOnlineByToken(ctx context.Context, token string) (err error) {
	_, err = dao.AdminUserOnline.Ctx(ctx).Delete(dao.AdminUserOnline.Columns().Token, token)
	return
}

func (s *sAdminUserOnline) ForceLogout(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var onlineList []*entity.AdminUserOnline
		onlineList, err = s.GetInfosByIds(ctx, ids)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.AdminUserOnline.Ctx(ctx).Where(dao.AdminUserOnline.Columns().Id+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err)
		for _, v := range onlineList {
			err = service.GfToken().RemoveToken(ctx, v.Token)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sAdminUserOnline) GetInfosByIds(ctx context.Context, ids []int) (onlineList []*entity.AdminUserOnline, err error) {
	err = dao.AdminUserOnline.Ctx(ctx).Where(dao.AdminUserOnline.Columns().Id+" in(?)", ids).Scan(&onlineList)
	return
}
