/*
* @desc:配置参数管理
* @company:
* @Author:
* @Date:   2022/9/28 9:13
 */

package adminConfig

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/consts"
	"github.com/mndon/gf-extensions/adminx/internal/dao"
	"github.com/mndon/gf-extensions/adminx/internal/lib/liberr"
	"github.com/mndon/gf-extensions/adminx/internal/model/do"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
	"github.com/mndon/gf-extensions/adminx/internal/service"
	"github.com/pkg/errors"
)

func init() {
	service.RegisterAdminConfig(New())
}

func New() *sAdminConfig {
	return &sAdminConfig{}
}

type sAdminConfig struct {
}

// List 系统参数列表
func (s *sAdminConfig) List(ctx context.Context, req *api.ConfigSearchReq) (res *api.ConfigSearchRes, err error) {
	res = new(api.ConfigSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AdminConfig.Ctx(ctx)
		if req != nil {
			if req.ConfigName != "" {
				m = m.Where("config_name like ?", "%"+req.ConfigName+"%")
			}
			if req.ConfigType != "" {
				m = m.Where("config_type = ", gconv.Int(req.ConfigType))
			}
			if req.ConfigKey != "" {
				m = m.Where("config_key like ?", "%"+req.ConfigKey+"%")
			}
			if len(req.DateRange) > 0 {
				m = m.Where("created_at >= ? AND created_at<=?", req.DateRange[0], req.DateRange[1])
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		err = m.Page(req.PageNum, req.PageSize).Order("config_id asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sAdminConfig) Add(ctx context.Context, req *api.ConfigAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.AdminConfig.Ctx(ctx).Insert(do.AdminConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			CreateBy:    userId,
			Remark:      req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加系统参数失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminConfigTag)
	})
	return
}

// CheckConfigKeyUnique 验证参数键名是否存在
func (s *sAdminConfig) CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := (*entity.AdminConfig)(nil)
		m := dao.AdminConfig.Ctx(ctx).Fields(dao.AdminConfig.Columns().ConfigId).Where(dao.AdminConfig.Columns().ConfigKey, configKey)
		if len(configId) > 0 {
			m = m.Where(dao.AdminConfig.Columns().ConfigId+" != ?", configId[0])
		}
		err = m.Scan(&data)
		liberr.ErrIsNil(ctx, err, "校验失败")
		if data != nil {
			liberr.ErrIsNil(ctx, errors.New("参数键名重复"))
		}
	})
	return
}

// Get 获取系统参数
func (s *sAdminConfig) Get(ctx context.Context, id int) (res *api.ConfigGetRes, err error) {
	res = new(api.ConfigGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminConfig.Ctx(ctx).WherePri(id).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "获取系统参数失败")
	})
	return
}

// Edit 修改系统参数
func (s *sAdminConfig) Edit(ctx context.Context, req *api.ConfigEditReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey, req.ConfigId)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.AdminConfig.Ctx(ctx).WherePri(req.ConfigId).Update(do.AdminConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			UpdateBy:    userId,
			Remark:      req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "修改系统参数失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminConfigTag)
	})
	return
}

// Delete 删除系统参数
func (s *sAdminConfig) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AdminConfig.Ctx(ctx).Delete(dao.AdminConfig.Columns().ConfigId+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
		//清除缓存
		service.Cache().RemoveByTag(ctx, consts.CacheAdminConfigTag)
	})
	return
}

// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sAdminConfig) GetConfigByKey(ctx context.Context, key string) (config *entity.AdminConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cache := service.Cache()
	cf := cache.Get(ctx, consts.CacheAdminConfigTag+key)
	if cf != nil && !cf.IsEmpty() {
		err = gconv.Struct(cf, &config)
		return
	}
	config, err = s.GetByKey(ctx, key)
	if err != nil {
		return
	}
	if config != nil {
		cache.Set(ctx, consts.CacheAdminConfigTag+key, config, 0, consts.CacheAdminConfigTag)
	}
	return
}

// GetByKey 通过key获取参数（从数据库获取）
func (s *sAdminConfig) GetByKey(ctx context.Context, key string) (config *entity.AdminConfig, err error) {
	err = dao.AdminConfig.Ctx(ctx).Where("config_key", key).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取配置失败")
	}
	return
}
