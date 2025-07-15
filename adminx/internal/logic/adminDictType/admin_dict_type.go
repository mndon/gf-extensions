/*
* @desc:字典类型管理
* @company:
* @Author:
* @Date:   2022/9/28 9:26
 */

package adminDictType

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
)

func init() {
	service.RegisterAdminDictType(New())
}

func New() *sAdminDictType {
	return &sAdminDictType{}
}

type sAdminDictType struct {
}

// List 字典类型列表
func (s *sAdminDictType) List(ctx context.Context, req *api.DictTypeSearchReq) (res *api.DictTypeSearchRes, err error) {
	res = new(api.DictTypeSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminDictType.Ctx(ctx)
		if req.DictName != "" {
			m = m.Where(dao.AdminDictType.Columns().DictName+" like ?", "%"+req.DictName+"%")
		}
		if req.DictType != "" {
			m = m.Where(dao.AdminDictType.Columns().DictType+" like ?", "%"+req.DictType+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.AdminDictType.Columns().Status+" = ", gconv.Int(req.Status))
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		err = m.Fields(model.AdminDictTypeInfoRes{}).Page(req.PageNum, req.PageSize).
			Order(dao.AdminDictType.Columns().DictId + " asc").Scan(&res.DictTypeList)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// Add 添加字典类型
func (s *sAdminDictType) Add(ctx context.Context, req *api.DictTypeAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.ExistsDictType(ctx, req.DictType)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.AdminDictType.Ctx(ctx).Insert(do.AdminDictType{
			DictName: req.DictName,
			DictType: req.DictType,
			Status:   req.Status,
			CreateBy: userId,
			Remark:   req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加字典类型失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
	})
	return
}

// Edit 修改字典类型
func (s *sAdminDictType) Edit(ctx context.Context, req *api.DictTypeEditReq, userId uint64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			err = s.ExistsDictType(ctx, req.DictType, req.DictId)
			liberr.ErrIsNil(ctx, err)
			dictType := (*entity.AdminDictType)(nil)
			e := dao.AdminDictType.Ctx(ctx).Fields(dao.AdminDictType.Columns().DictType).WherePri(req.DictId).Scan(&dictType)
			liberr.ErrIsNil(ctx, e, "获取字典类型失败")
			liberr.ValueIsNil(dictType, "字典类型不存在")
			//修改字典类型
			_, e = dao.AdminDictType.Ctx(ctx).TX(tx).WherePri(req.DictId).Update(do.AdminDictType{
				DictName: req.DictName,
				DictType: req.DictType,
				Status:   req.Status,
				UpdateBy: userId,
				Remark:   req.Remark,
			})
			liberr.ErrIsNil(ctx, e, "修改字典类型失败")
			//修改字典数据
			_, e = dao.AdminDictData.Ctx(ctx).TX(tx).Data(do.AdminDictData{DictType: req.DictType}).
				Where(dao.AdminDictData.Columns().DictType, dictType.DictType).Update()
			liberr.ErrIsNil(ctx, e, "修改字典数据失败")
			//清除缓存
			service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
		})
		return err
	})
	return
}

func (s *sAdminDictType) Get(ctx context.Context, req *api.DictTypeGetReq) (dictType *entity.AdminDictType, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminDictType.Ctx(ctx).Where(dao.AdminDictType.Columns().DictId, req.DictId).Scan(&dictType)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// ExistsDictType 检查类型是否已经存在
func (s *sAdminDictType) ExistsDictType(ctx context.Context, dictType string, dictId ...int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminDictType.Ctx(ctx).Fields(dao.AdminDictType.Columns().DictId).
			Where(dao.AdminDictType.Columns().DictType, dictType)
		if len(dictId) > 0 {
			m = m.Where(dao.AdminDictType.Columns().DictId+" !=? ", dictId[0])
		}
		res, e := m.One()
		liberr.ErrIsNil(ctx, e, "sql err")
		if !res.IsEmpty() {
			liberr.ErrIsNil(ctx, gerror.New("字典类型已存在"))
		}
	})
	return
}

// Delete 删除字典类型
func (s *sAdminDictType) Delete(ctx context.Context, dictIds []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			discs := ([]*entity.AdminDictType)(nil)
			err = dao.AdminDictType.Ctx(ctx).Fields(dao.AdminDictType.Columns().DictType).
				Where(dao.AdminDictType.Columns().DictId+" in (?) ", dictIds).Scan(&discs)
			liberr.ErrIsNil(ctx, err, "删除失败")
			types := garray.NewStrArray()
			for _, dt := range discs {
				types.Append(dt.DictType)
			}
			if types.Len() > 0 {
				_, err = dao.AdminDictType.Ctx(ctx).TX(tx).Delete(dao.AdminDictType.Columns().DictId+" in (?) ", dictIds)
				liberr.ErrIsNil(ctx, err, "删除类型失败")
				_, err = dao.AdminDictData.Ctx(ctx).TX(tx).Delete(dao.AdminDictData.Columns().DictType+" in (?) ", types.Slice())
				liberr.ErrIsNil(ctx, err, "删除字典数据失败")
			}
			//清除缓存
			service.Cache().RemoveByTag(ctx, consts.CacheAdminDictTag)
		})
		return err
	})
	return
}

// GetAllDictType 获取所有正常状态下的字典类型
func (s *sAdminDictType) GetAllDictType(ctx context.Context) (list []*entity.AdminDictType, err error) {
	cache := service.Cache()
	//从缓存获取
	data := cache.Get(ctx, consts.CacheAdminDict+"_dict_type_all")
	if !data.IsNil() {
		err = data.Structs(&list)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminDictType.Ctx(ctx).Where("status", 1).Order("dict_id ASC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取字典类型数据出错")
		//缓存
		cache.Set(ctx, consts.CacheAdminDict+"_dict_type_all", list, 0, consts.CacheAdminDictTag)
	})
	return
}
