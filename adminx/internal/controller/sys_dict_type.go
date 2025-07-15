/*
* @desc:字典类型
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

var DictType = &AdminDictTypeController{}

type AdminDictTypeController struct {
}

// List 字典类型列表
func (c *AdminDictTypeController) List(ctx context.Context, req *api.DictTypeSearchReq) (res *api.DictTypeSearchRes, err error) {
	res, err = service.AdminDictType().List(ctx, req)
	return
}

// Add 添加字典类型
func (c *AdminDictTypeController) Add(ctx context.Context, req *api.DictTypeAddReq) (res *api.DictTypeAddRes, err error) {
	err = service.AdminDictType().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取字典类型
func (c *AdminDictTypeController) Get(ctx context.Context, req *api.DictTypeGetReq) (res *api.DictTypeGetRes, err error) {
	res = new(api.DictTypeGetRes)
	res.DictType, err = service.AdminDictType().Get(ctx, req)
	return
}

// Edit 修改字典数据
func (c *AdminDictTypeController) Edit(ctx context.Context, req *api.DictTypeEditReq) (res *api.DictTypeEditRes, err error) {
	err = service.AdminDictType().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

func (c *AdminDictTypeController) Delete(ctx context.Context, req *api.DictTypeDeleteReq) (res *api.DictTypeDeleteRes, err error) {
	err = service.AdminDictType().Delete(ctx, req.DictIds)
	return
}

// OptionSelect 获取字典选择框列表
func (c *AdminDictTypeController) OptionSelect(ctx context.Context, req *api.DictTypeAllReq) (res *api.DictTYpeAllRes, err error) {
	res = new(api.DictTYpeAllRes)
	res.DictType, err = service.AdminDictType().GetAllDictType(ctx)
	return
}
