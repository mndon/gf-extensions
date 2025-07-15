/*
* @desc:字典数据管理
* @company:
* @Author:
* @Date:   2022/9/28 9:22
 */

package adminDictData

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterAdminDictData(New())
}

func New() *sAdminDictData {
	return &sAdminDictData{}
}

type sAdminDictData struct {
}

// GetDictWithDataByType 通过字典键类型获取选项
func (s *sAdminDictData) GetDictWithDataByType(ctx context.Context, dictType, defaultValue string) (dict *api.GetDictRes,
	err error) {
	cache := service.Cache()
	cacheKey := consts.CacheAdminDict + "_" + dictType
	//从缓存获取
	iDict := cache.GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (value interface{}, err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			//从数据库获取
			dict = &api.GetDictRes{}
			//获取类型数据
			err = dao.AdminDictType.Ctx(ctx).Where(dao.AdminDictType.Columns().DictType, dictType).
				Where(dao.AdminDictType.Columns().Status, 1).Fields(model.DictTypeRes{}).Scan(&dict.Info)
			liberr.ErrIsNil(ctx, err, "获取字典类型失败")
			if dict.Info == nil {
				return
			}
			err = dao.AdminDictData.Ctx(ctx).Fields(model.DictDataRes{}).
				Where(dao.AdminDictData.Columns().DictType, dictType).
				Where(dao.AdminDictData.Columns().Status, 1).
				Order(dao.AdminDictData.Columns().DictSort + " asc," +
					dao.AdminDictData.Columns().DictCode + " asc").
				Scan(&dict.Values)
			liberr.ErrIsNil(ctx, err, "获取字典数据失败")
		})
		value = dict
		return
	}, 0, consts.CacheAdminDictTag)
	if !iDict.IsEmpty() {
		err = gconv.Struct(iDict, &dict)
		if err != nil {
			return
		}
	}
	//设置给定的默认值
	for _, v := range dict.Values {
		if defaultValue != "" {
			if gstr.Equal(defaultValue, v.DictValue) {
				v.IsDefault = 1
			} else {
				v.IsDefault = 0
			}
		}
	}
	return
}

// List 获取字典数据
func (s *sAdminDictData) List(ctx context.Context, req *api.DictDataSearchReq) (res *api.DictDataSearchRes, err error) {
	res = new(api.DictDataSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminDictData.Ctx(ctx)
		if req != nil {
			if req.DictLabel != "" {
				m = m.Where(dao.AdminDictData.Columns().DictLabel+" like ?", "%"+req.DictLabel+"%")
			}
			if req.Status != "" {
				m = m.Where(dao.AdminDictData.Columns().Status+" = ", gconv.Int(req.Status))
			}
			if req.DictType != "" {
				m = m.Where(dao.AdminDictData.Columns().DictType+" = ?", req.DictType)
			}
			res.Total, err = m.Count()
			liberr.ErrIsNil(ctx, err, "获取字典数据失败")
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			res.CurrentPage = req.PageNum
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		err = m.Page(req.PageNum, req.PageSize).Order(dao.AdminDictData.Columns().DictSort + " asc," +
			dao.AdminDictData.Columns().DictCode + " asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取字典数据失败")
	})
	return
}

func (s *sAdminDictData) Add(ctx context.Context, req *api.DictDataAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminDictData.Ctx(ctx).Insert(do.AdminDictData{
			DictSort:  req.DictSort,
			DictLabel: req.DictLabel,
			DictValue: req.DictValue,
			DictType:  req.DictType,
			CssClass:  req.CssClass,
			ListClass: req.ListClass,
			IsDefault: req.IsDefault,
			Status:    req.Status,
			CreateBy:  userId,
			Remark:    req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加字典数据失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
	})
	return
}

// Get 获取字典数据
func (s *sAdminDictData) Get(ctx context.Context, dictCode uint) (res *api.DictDataGetRes, err error) {
	res = new(api.DictDataGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminDictData.Ctx(ctx).WherePri(dictCode).Scan(&res.Dict)
		liberr.ErrIsNil(ctx, err, "获取字典数据失败")
	})
	return
}

// Edit 修改字典数据
func (s *sAdminDictData) Edit(ctx context.Context, req *api.DictDataEditReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminDictData.Ctx(ctx).WherePri(req.DictCode).Update(do.AdminDictData{
			DictSort:  req.DictSort,
			DictLabel: req.DictLabel,
			DictValue: req.DictValue,
			DictType:  req.DictType,
			CssClass:  req.CssClass,
			ListClass: req.ListClass,
			IsDefault: req.IsDefault,
			Status:    req.Status,
			UpdateBy:  userId,
			Remark:    req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "修改字典数据失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
	})
	return
}

// Delete 删除字典数据
func (s *sAdminDictData) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminDictData.Ctx(ctx).Where(dao.AdminDictData.Columns().DictCode+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除字典数据失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
	})
	return
}
