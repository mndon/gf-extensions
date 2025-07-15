/*
* @desc:公用model
* @company:
* @Author:
* @Date:   2023/5/11 22:43
 */

package model

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `p:"dateRange"`                                                                                                                                                //日期范围
	PageNum   int      `p:"pageNum"`                                                                                                                                                  //当前页码
	PageSize  int      `p:"pageSize"`                                                                                                                                                 //每页数
	OrderBy   string   `p:"orderBy" v:"regex:^[a-zA-Z0-9_]+(\\.[a-zA-Z0-9_]+)?\\s+(asc|desc|ASC|DESC)(?:\\s*,\\s*[a-zA-Z0-9_]+(\\.[a-zA-Z0-9_]+)?\\s+(asc|desc|ASC|DESC))*$#排序参数不合法"` // 排序方式
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
