package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/kvconfigx"
)

type KvConfigListReq struct {
	g.Meta `path:"/api/v1/kv_config/list" method:"get"`
	model.KvConfigListInput
}

type KvConfigListRes struct {
	model.KvConfigListOutput
}

type KvConfigGetReq struct {
	g.Meta `path:"/api/v1/kv_config/get" method:"get"`
	K      string `json:"k"`
}

type KvConfigGetRes struct {
	*kvconfigx.KvConfigEntity
}

type KvConfigCreateReq struct {
	g.Meta `path:"/api/v1/kv_config/create" method:"post"`
	kvconfigx.KvConfigDo
}

type KvConfigCreateRes struct {
	Id int64 `json:"id"`
}

type KvConfigUpdateReq struct {
	g.Meta `path:"/api/v1/kv_config/update" method:"post"`
	kvconfigx.KvConfigDo
}

type KvConfigUpdateRes struct {
}

type KvConfigDeleteReq struct {
	g.Meta `path:"/api/v1/kv_config/delete" method:"post"`
}

type KvConfigDeleteRes struct {
}
