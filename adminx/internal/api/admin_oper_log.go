/*
* @desc:操作日志
* @company:
* @Author:
* @Date:   2022/12/21 14:37
 */

package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/adminx/internal/model"
)

// AdminOperLogSearchReq 分页请求参数
type AdminOperLogSearchReq struct {
	g.Meta        `path:"/operLog/list" tags:"操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`         //系统模块
	RequestMethod string `p:"requestMethod"` //请求方式
	OperName      string `p:"operName"`      //操作人员
	PageReq
	Author
}

// AdminOperLogSearchRes 列表返回结果
type AdminOperLogSearchRes struct {
	g.Meta `mime:"application/json"`
	ListRes
	List []*model.AdminOperLogListRes `json:"list"`
}

// AdminOperLogGetReq 获取一条数据请求
type AdminOperLogGetReq struct {
	g.Meta `path:"/operLog/get" tags:"操作日志" method:"get" summary:"获取操作日志信息"`
	Author
	OperId uint64 `p:"operId" v:"required#主键必须"` //通过主键获取
}

// AdminOperLogGetRes 获取一条数据结果
type AdminOperLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.AdminOperLogInfoRes
}

// AdminOperLogDeleteReq 删除数据请求
type AdminOperLogDeleteReq struct {
	g.Meta `path:"/operLog/delete" tags:"操作日志" method:"delete" summary:"删除操作日志"`
	Author
	OperIds []uint64 `p:"operIds" v:"required#主键必须"` //通过主键删除
}

// AdminOperLogDeleteRes 删除数据返回
type AdminOperLogDeleteRes struct {
	EmptyRes
}

type AdminOperLogClearReq struct {
	g.Meta `path:"/operLog/clear" tags:"操作日志" method:"delete" summary:"清除日志"`
	Author
}

type AdminOperLogClearRes struct {
	EmptyRes
}
