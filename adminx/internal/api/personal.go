/*
* @desc:xxxx功能描述
* @company:
* @Author:
* @Date:   2022/11/3 10:04
 */

package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mndon/gf-extensions/adminx/internal/model"
	"github.com/mndon/gf-extensions/adminx/internal/model/entity"
)

type PersonalInfoReq struct {
	g.Meta `path:"/personal/getPersonalInfo" tags:"用户管理" method:"get" summary:"登录用户信息"`
	Author
}

type PersonalInfoRes struct {
	g.Meta   `mime:"application/json"`
	User     *entity.AdminUser `json:"user"`
	Roles    []string          `json:"roles"`
	DeptName string            `json:"deptName"`
}

// SetPersonalReq 添加修改用户公用请求字段
type SetPersonalReq struct {
	Nickname  string `p:"nickname" v:"required#用户昵称不能为空"`
	Mobile    string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Remark    string `p:"remark"`
	Sex       int    `p:"sex"`
	UserEmail string `p:"userEmail" v:"required|email#邮箱不能为空|邮箱格式错误"`
	Describe  string `p:"describe"` //签名
	Avatar    string `p:"avatar"`   //签名

}

// PersonalEditReq 修改个人
type PersonalEditReq struct {
	g.Meta `path:"/personal/edit" tags:"用户管理" method:"put" summary:"修改个人资料"`
	*SetPersonalReq
	Author
}

type PersonalEditRes struct {
	EmptyRes
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type PersonalResetPwdReq struct {
	g.Meta   `path:"/personal/resetPwd" tags:"用户管理" method:"put" summary:"重置个人密码"`
	Password string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Author
}

type PersonalResetPwdRes struct {
}
