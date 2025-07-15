/*
* @desc:登录日志
* @company:
* @Author:
* @Date:   2022/9/26 15:20
 */

package adminLoginLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterAdminLoginLog(New())
}

func New() *sAdminLoginLog {
	return &sAdminLoginLog{
		Pool: grpool.New(100),
	}
}

type sAdminLoginLog struct {
	Pool *grpool.Pool
}

func (s *sAdminLoginLog) Invoke(ctx context.Context, data *model.LoginLogParams) {
	s.Pool.Add(
		ctx,
		func(ctx context.Context) {
			//写入日志数据
			service.AdminUser().LoginLog(ctx, data)
		},
	)
}

func (s *sAdminLoginLog) List(ctx context.Context, req *api.LoginLogSearchReq) (res *api.LoginLogSearchRes, err error) {
	res = new(api.LoginLogSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	m := dao.AdminLoginLog.Ctx(ctx)
	order := "info_id DESC"
	if req.LoginName != "" {
		m = m.Where("login_name like ?", "%"+req.LoginName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", gconv.Int(req.Status))
	}
	if req.Ipaddr != "" {
		m = m.Where("ipaddr like ?", "%"+req.Ipaddr+"%")
	}
	if req.LoginLocation != "" {
		m = m.Where("login_location like ?", "%"+req.LoginLocation+"%")
	}
	if len(req.DateRange) != 0 {
		m = m.Where("login_time >=? AND login_time <=?", req.DateRange[0], req.DateRange[1])
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取日志失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取日志数据失败")
	})
	return
}

func (s *sAdminLoginLog) DeleteLoginLogByIds(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminLoginLog.Ctx(ctx).Delete("info_id in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sAdminLoginLog) ClearLoginLog(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = g.DB().Ctx(ctx).Exec(ctx, "truncate "+dao.AdminLoginLog.Table())
		liberr.ErrIsNil(ctx, err, "清除失败")
	})
	return
}
