/*
* @desc:字典数据管理
* @company:
* @Author:
* @Date:   2022/3/18 11:57
 */

package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

var DictData = dictDataController{}

type dictDataController struct {
}

// GetDictData 获取字典数据
func (c *dictDataController) GetDictData(ctx context.Context, req *api.GetDictReq) (res *api.GetDictRes, err error) {
	res, err = service.AdminDictData().GetDictWithDataByType(ctx, req.DictType, req.DefaultValue)
	return
}

// List 获取字典数据列表
func (c *dictDataController) List(ctx context.Context, req *api.DictDataSearchReq) (res *api.DictDataSearchRes, err error) {
	res, err = service.AdminDictData().List(ctx, req)
	return
}

// Add 添加字典数据
func (c *dictDataController) Add(ctx context.Context, req *api.DictDataAddReq) (res *api.DictDataAddRes, err error) {
	err = service.AdminDictData().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取对应的字典数据
func (c *dictDataController) Get(ctx context.Context, req *api.DictDataGetReq) (res *api.DictDataGetRes, err error) {
	res, err = service.AdminDictData().Get(ctx, req.DictCode)
	return
}

// Edit 修改字典数据
func (c *dictDataController) Edit(ctx context.Context, req *api.DictDataEditReq) (res *api.DictDataEditRes, err error) {
	err = service.AdminDictData().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

func (c *dictDataController) Delete(ctx context.Context, req *api.DictDataDeleteReq) (res *api.DictDataDeleteRes, err error) {
	err = service.AdminDictData().Delete(ctx, req.Ids)
	return
}
