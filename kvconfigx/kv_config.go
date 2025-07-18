package kvconfigx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mndon/gf-extensions/errorx"
	"github.com/mndon/gf-extensions/kvconfigx/internal/dao"
	"github.com/mndon/gf-extensions/kvconfigx/internal/model/do"
	"github.com/mndon/gf-extensions/kvconfigx/internal/model/entity"
	"strings"
)

type KvConfigEntity = entity.KvConfig
type KvConfigDo = do.KvConfig

// Get
// @Description: 获取key配置
// @param ctx
// @param key
// @return out
// @return err
func Get(ctx context.Context, key string) (out *gvar.Var, err error) {
	var item *entity.KvConfig
	err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errorx.NotFoundErr("no.such.config", "无此配置，请联系客服解决")
	}

	return gvar.New(item.V), nil
}

// GetAndFormatToMapOrArr
// @Description: 将
// @param ctx
// @param v
// @return out
// @return err
func GetAndFormatToMapOrArr(ctx context.Context, key string) (out any, err error) {
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
	var item *entity.KvConfig
	err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Scan(&item)
	if err != nil {
		return err
	}
	if item == nil {
		return nil
	}

	err = gconv.Scan(item.V, pointer)
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

	var item *entity.KvConfig
	err = dao.KvConfig.Ctx(ctx).Where(do.KvConfig{K: key}).Scan(&item)
	if err != nil {
		return false, err
	}
	if item == nil {
		return false, errorx.NotFoundErr("no.such.config", "无此配置，请联系客服解决")
	}
	return gconv.Bool(item.V), nil
}
