package kvconfigx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/errorx"
	"github.com/mndon/gf-extensions/kvconfigx/internal/dao"
	"github.com/mndon/gf-extensions/kvconfigx/internal/model/do"
	"github.com/mndon/gf-extensions/kvconfigx/internal/model/entity"
	"strings"
	"time"
)

type KvConfigEntity = entity.KvConfig
type KvConfigDo = do.KvConfig

// 本地缓存
var cache = gcache.New()

// Get
// @Description: 获取key配置
// @param ctx
// @param key
// @return out
// @return err
func Get(ctx context.Context, key string) (out *gvar.Var, err error) {
	// 读缓存
	cacheValue, err := cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if cacheValue != nil {
		return cacheValue, nil
	}

	// 读数据库，并缓存
	var item *entity.KvConfig
	err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errorx.NotFoundErr("no.such.config", "无此配置，请联系客服解决")
	}
	err = cache.Set(ctx, key, item.V, time.Second*10)
	if err != nil {
		return nil, err
	}

	return gvar.New(item.V), nil
}

// GetAndFormatToMapOrSlices
// @Description: 将
// @param ctx
// @param v
// @return out
// @return err
func GetAndFormatToMapOrSlices(ctx context.Context, key string) (out any, err error) {
	v, err := Get(ctx, key)
	if err != nil {
		return nil, err
	}

	vStr := v.String()

	if strings.HasPrefix(vStr, "[") {
		var result []any
		err = gconv.Scan(vStr, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		var result map[string]interface{}
		err = gconv.Scan(vStr, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

// Scan
// @Description: 获取key配置
// @param ctx
// @param key
// @param pointer
// @return err
func Scan(ctx context.Context, key string, pointer any) (err error) {
	vVar, err := Get(ctx, key)
	if err != nil {
		return err
	}

	err = vVar.Scan(pointer)
	if err != nil {
		return err
	}
	return nil
}

// GetBool
// @Description: 获取key配置
// @param ctx
// @param key
// @return out
// @return err
func GetBool(ctx context.Context, key string, def ...bool) (out bool, err error) {
	defer func() {
		if err != nil && len(def) > 0 {
			out = def[0]
			err = nil
		}
	}()

	cacheValue, err := Get(ctx, key)
	if err != nil {
		return false, err
	}

	return cacheValue.Bool(), nil
}

// ------------ 以下为后台方法 ------------

// AdminCreate
// @Description: 新增
// @param ctx
// @param data
// @return err
func AdminCreate(ctx context.Context, data KvConfigDo) (err error) {
	_, err = dao.KvConfig.Ctx(ctx).Data(data).Insert()
	return err
}

// AdminUpdate
// @Description: 更新key配置
// @param ctx
// @param key
// @param value
// @return err
func AdminUpdate(ctx context.Context, data KvConfigDo) (err error) {
	data.Id = nil
	data.K = nil
	data.UpdatedTime = nil
	data.CreatedTime = nil
	_, err = dao.KvConfig.Ctx(ctx).Data(data).Where(do.KvConfig{K: data.K}).Update()
	if err != nil {
		return err
	}

	// 移除缓存
	_, err = cache.Remove(ctx, data.K)
	return err
}

// AdminDelete
// @Description: 删除key配置
// @param ctx
// @param key
// @return err
func AdminDelete(ctx context.Context, key string) (err error) {
	_, err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Delete()
	if err != nil {
		return err
	}

	// 移除缓存
	_, err = cache.Remove(ctx, key)
	return err
}

// AdminGet
// @Description: 获取key配置
// @param ctx
// @param key
// @return item
// @return err
func AdminGet(ctx context.Context, key string) (item *entity.KvConfig, err error) {
	err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Scan(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// AdminList
// @Description: 获取kv配置列表
// @param ctx
// @param key
// @param page
// @param size
// @return list
// @return total
// @return err
func AdminList(ctx context.Context, key string, page, size int) (list []entity.KvConfig, total int, err error) {
	conn := dao.KvConfig.Ctx(ctx)

	if key != "" {
		conn = conn.WhereLike(dao.KvConfig.Columns().K, key)
	}
	total, err = conn.Count()
	if err != nil {
		return nil, 0, err
	}

	err = conn.Page(page, size).OrderDesc(dao.KvConfig.Columns().Id).Scan(list)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
