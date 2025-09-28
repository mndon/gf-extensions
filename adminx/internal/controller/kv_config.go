package controller

import (
	"context"
	"github.com/mndon/gf-extensions/adminx/internal/api"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/kvconfigx"
)

type kvConfig struct {
}

var KvConfig = kvConfig{}

func (k *kvConfig) List(ctx context.Context, req *api.KvConfigListReq) (res *api.KvConfigListRes, err error) {
	list, total, err := kvconfigx.AdminList(ctx, req.K, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &api.KvConfigListRes{
		KvConfigListOutput: model.KvConfigListOutput{
			List:  list,
			Page:  req.Page,
			Size:  req.Size,
			Total: total,
		},
	}, nil
}

func (k *kvConfig) Get(ctx context.Context, req *api.KvConfigGetReq) (res *api.KvConfigGetRes, err error) {
	out, err := kvconfigx.AdminGet(ctx, req.K)
	if err != nil {
		return nil, err
	}
	return &api.KvConfigGetRes{
		KvConfigEntity: out,
	}, nil
}

func (k *kvConfig) Update(ctx context.Context, req *api.KvConfigUpdateReq) (res *api.KvConfigUpdateRes, err error) {
	err = kvconfigx.AdminUpdate(ctx, req.KvConfigDo)
	if err != nil {
		return nil, err
	}
	return
}

func (k *kvConfig) Create(ctx context.Context, req *api.KvConfigCreateReq) (res *api.KvConfigCreateRes, err error) {
	err = kvconfigx.AdminCreate(ctx, req.KvConfigDo)
	if err != nil {
		return nil, err
	}
	return
}
