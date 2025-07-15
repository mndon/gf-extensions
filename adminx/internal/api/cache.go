/*
* @desc:缓存处理
* @company:
* @Author:
* @Date:   2023/2/1 18:12
 */

package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"缓存管理" method:"delete" summary:"清除缓存"`
	Author
}

type CacheRemoveRes struct {
	EmptyRes
}
