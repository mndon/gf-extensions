package uidx

import (
	"github.com/yitter/idgenerator-go/idgen"
)

// Init 单例初始化
func Init(workId uint16) {
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(workId))
}

// NextId
// @Description: 获取雪花id
// @return int64
func NextId() int64 {
	return idgen.NextId()
}
