package configx

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Init 修改配置为Adapter，默认读取cmd参数+配置文件, 根据环境区分配置文件
func Init() {
	g.Cfg().SetAdapter(NewAdapterFileWithCmd())
}
