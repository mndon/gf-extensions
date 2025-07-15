/*
* @desc:公共接口相关
* @company:
* @Author:
* @Date:   2022/3/30 9:28
 */

package api

import "github.com/mndon/gf-extensions/adminx/internal/model"

// PageReq 公共请求参数
type PageReq struct {
	model.PageReq
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
