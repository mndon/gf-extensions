/*
* @desc:系统后台操作日志
* @company:
* @Author:
* @Date:   2022/9/21 16:10
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var OperLog = new(operateLogController)

type operateLogController struct {
	BaseController
}

// List 列表
func (c *operateLogController) List(ctx context.Context, req *api.AdminOperLogSearchReq) (res *api.AdminOperLogSearchRes, err error) {
	res, err = service.OperateLog().List(ctx, req)
	return
}

// Get 获取操作日志
func (c *operateLogController) Get(ctx context.Context, req *api.AdminOperLogGetReq) (res *api.AdminOperLogGetRes, err error) {
	res = new(api.AdminOperLogGetRes)
	res.AdminOperLogInfoRes, err = service.OperateLog().GetByOperId(ctx, req.OperId)
	return
}

func (c *operateLogController) Delete(ctx context.Context, req *api.AdminOperLogDeleteReq) (res *api.AdminOperLogDeleteRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperIds)
	return
}

func (c *operateLogController) Clear(ctx context.Context, req *api.AdminOperLogClearReq) (res *api.AdminOperLogClearRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
