/*
* @desc:后台操作日志业务处理
* @company:
* @Author:
* @Date:   2022/9/21 16:14
 */

package adminOperLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/libUtils"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

type sOperateLog struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterOperateLog(New())
}

func New() *sOperateLog {
	return &sOperateLog{
		Pool: grpool.New(100),
	}
}

// OperationLog 操作日志写入
func (s *sOperateLog) OperationLog(r *ghttp.Request) {
	userInfo := service.Context().GetLoginUser(r.GetCtx())
	if userInfo == nil {
		return
	}
	url := r.Request.URL //请求地址
	//获取菜单
	//获取地址对应的菜单id
	menuList, err := service.AdminAuthRule().GetMenuList(r.GetCtx())
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	var menu *model.AdminAuthRuleInfoRes
	path := gstr.TrimLeft(url.Path, "/")
	for _, m := range menuList {
		if gstr.Equal(m.Name, path) {
			menu = m
			break
		}
	}
	data := &model.AdminOperLogAdd{
		User:         userInfo,
		Menu:         menu,
		Url:          url,
		Params:       r.GetMap(),
		Method:       r.Method,
		ClientIp:     libUtils.GetClientIp(r.GetCtx()),
		OperatorType: 1,
	}
	s.Invoke(gctx.New(), data)
}

func (s *sOperateLog) Invoke(ctx context.Context, data *model.AdminOperLogAdd) {
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		s.operationLogAdd(ctx, data)
	})
}

// OperationLogAdd 添加操作日志
func (s *sOperateLog) operationLogAdd(ctx context.Context, data *model.AdminOperLogAdd) {
	menuTitle := ""
	if data.Menu != nil {
		menuTitle = data.Menu.Title
	}
	insertData := &do.AdminOperLog{
		Title:         menuTitle,
		Method:        data.Url.Path,
		RequestMethod: data.Method,
		OperatorType:  data.OperatorType,
		OperName:      data.User.UserName,
		OperIp:        data.ClientIp,
		OperLocation:  libUtils.GetCityByIp(data.ClientIp),
		OperTime:      gtime.Now(),
		OperParam:     data.Params,
	}
	rawQuery := data.Url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	insertData.OperUrl = data.Url.Path + rawQuery
	_, err := dao.AdminOperLog.Ctx(ctx).Insert(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sOperateLog) List(ctx context.Context, req *api.AdminOperLogSearchReq) (listRes *api.AdminOperLogSearchRes, err error) {
	listRes = new(api.AdminOperLogSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminOperLog.Ctx(ctx)
		if req.Title != "" {
			m = m.Where(dao.AdminOperLog.Columns().Title+" = ?", req.Title)
		}
		if req.RequestMethod != "" {
			m = m.Where(dao.AdminOperLog.Columns().RequestMethod+" = ?", req.RequestMethod)
		}
		if req.OperName != "" {
			m = m.Where(dao.AdminOperLog.Columns().OperName+" like ?", "%"+req.OperName+"%")
		}
		if len(req.DateRange) != 0 {
			m = m.Where("oper_time >=? AND oper_time <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		order := "oper_id DESC"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.AdminOperLogInfoRes
		err = m.Fields(api.AdminOperLogSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.AdminOperLogListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.AdminOperLogListRes{
				OperId:         v.OperId,
				Title:          v.Title,
				RequestMethod:  v.RequestMethod,
				OperName:       v.OperName,
				DeptName:       v.DeptName,
				LinkedDeptName: v.LinkedDeptName,
				OperUrl:        v.OperUrl,
				OperIp:         v.OperIp,
				OperLocation:   v.OperLocation,
				OperParam:      v.OperParam,
				OperTime:       v.OperTime,
			}
		}
	})
	return
}

func (s *sOperateLog) GetByOperId(ctx context.Context, operId uint64) (res *model.AdminOperLogInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminOperLog.Ctx(ctx).WithAll().Where(dao.AdminOperLog.Columns().OperId, operId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sOperateLog) DeleteByIds(ctx context.Context, ids []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminOperLog.Ctx(ctx).Delete("oper_id in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sOperateLog) ClearLog(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = g.DB().Ctx(ctx).Exec(ctx, "truncate "+dao.AdminOperLog.Table())
		liberr.ErrIsNil(ctx, err, "清除失败")
	})
	return
}
