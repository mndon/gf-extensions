package configx

import "github.com/gogf/gf/v2/frame/g"

// init
// @Description: 修改配置为Adapter，默认读取cmd参数+配置文件
func init() {
	g.Cfg().SetAdapter(NewAdapterFileWithCmd())
}
