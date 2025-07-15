/*
* @desc:返回响应公共参数
* @company:
* @Author:
* @Date:   2022/10/27 16:30
 */

package api

import "github.com/gogf/gf/v2/frame/g"

// EmptyRes 不响应任何数据
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
