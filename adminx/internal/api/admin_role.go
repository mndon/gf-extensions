/*
* @desc:角色api
* @company:
* @Author:
* @Date:   2022/3/30 9:16
 */

package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
)

type RoleListReq struct {
	g.Meta   `path:"/role/list" tags:"角色管理" method:"get" summary:"角色列表"`
	RoleName string `p:"roleName"`   //参数名称
	Status   string `p:"roleStatus"` //状态
	PageReq
}

type RoleListRes struct {
	g.Meta `mime:"application/json"`
	ListRes
	List []*entity.AdminRole `json:"list"`
}

type RoleGetParamsReq struct {
	g.Meta `path:"/role/getParams" tags:"角色管理" method:"get" summary:"角色编辑参数"`
}

type RoleGetParamsRes struct {
	g.Meta `mime:"application/json"`
	Menu   []*model.AdminAuthRuleInfoRes `json:"menu"`
}

type RoleAddReq struct {
	g.Meta    `path:"/role/add" tags:"角色管理" method:"post" summary:"添加角色"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
}

type RoleAddRes struct {
}

type RoleGetReq struct {
	g.Meta `path:"/role/get" tags:"角色管理" method:"get" summary:"获取角色信息"`
	Id     uint `p:"id" v:"required#角色id不能为空"`
}

type RoleGetRes struct {
	g.Meta  `mime:"application/json"`
	Role    *entity.AdminRole `json:"role"`
	MenuIds []int             `json:"menuIds"`
}

type RoleEditReq struct {
	g.Meta    `path:"/role/edit" tags:"角色管理" method:"put" summary:"修改角色"`
	Id        int64  `p:"id" v:"required#角色id必须"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
}

type RoleEditRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"角色管理" method:"delete" summary:"删除角色"`
	Ids    []int64 `p:"ids" v:"required#角色id不能为空"`
}

type RoleDeleteRes struct {
}
