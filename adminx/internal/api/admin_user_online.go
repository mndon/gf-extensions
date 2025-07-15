/*
* @desc:在线用户
* @company:
* @Author:
* @Date:   2023/1/10 16:57
 */

package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
)

// AdminUserOnlineSearchReq 列表搜索参数
type AdminUserOnlineSearchReq struct {
	g.Meta   `path:"/online/list" tags:"在线用户管理" method:"get" summary:"列表"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	PageReq
	Author
}

// AdminUserOnlineSearchRes 列表结果
type AdminUserOnlineSearchRes struct {
	g.Meta `mime:"application/json"`
	ListRes
	List []*entity.AdminUserOnline `json:"list"`
}

type AdminUserOnlineForceLogoutReq struct {
	g.Meta `path:"/online/forceLogout" tags:"在线用户管理" method:"delete" summary:"强制用户退出登录"`
	Author
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type AdminUserOnlineForceLogoutRes struct {
	EmptyRes
}
