package uidx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
	"sync"
)

var initOnce sync.Once

// NextId
// @Description: 获取雪花id
// @return int64
func NextId() int64 {
	initOnce.Do(func() {
		workId := uint16(0)
		v, _ := g.Cfg().Get(context.Background(), "uidx.workId", 0)
		if v != nil {
			workId = v.Uint16()
		}
		idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(workId))
	})
	return idgen.NextId()
}
