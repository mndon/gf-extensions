/*
* @desc:操作日志模型对象
* @company:
* @Author:
* @Date:   2022/9/21 16:34
 */

package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"net/url"
)

// AdminOperLogAdd 添加操作日志参数
type AdminOperLogAdd struct {
	User         *ContextUser
	Menu         *AdminAuthRuleInfoRes
	Url          *url.URL
	Params       g.Map
	Method       string
	ClientIp     string
	OperatorType int
}

// AdminOperLogInfoRes is the golang structure for table admin_oper_log.
type AdminOperLogInfoRes struct {
	gmeta.Meta     `orm:"table:admin_oper_log"`
	OperId         uint64                       `orm:"oper_id,primary" json:"operId"`       // 日志编号
	Title          string                       `orm:"title" json:"title"`                  // 系统模块
	BusinessType   int                          `orm:"business_type" json:"businessType"`   // 操作类型
	Method         string                       `orm:"method" json:"method"`                // 操作方法
	RequestMethod  string                       `orm:"request_method" json:"requestMethod"` // 请求方式
	OperatorType   int                          `orm:"operator_type" json:"operatorType"`   // 操作类别
	OperName       string                       `orm:"oper_name" json:"operName"`           // 操作人员
	DeptName       string                       `orm:"dept_name" json:"deptName"`           // 部门名称
	LinkedDeptName *LinkedAdminOperLogAdminDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperUrl        string                       `orm:"oper_url" json:"operUrl"`           // 请求URL
	OperIp         string                       `orm:"oper_ip" json:"operIp"`             // 主机地址
	OperLocation   string                       `orm:"oper_location" json:"operLocation"` // 操作地点
	OperParam      string                       `orm:"oper_param" json:"operParam"`       // 请求参数
	ErrorMsg       string                       `orm:"error_msg" json:"errorMsg"`         // 错误消息
	OperTime       *gtime.Time                  `orm:"oper_time" json:"operTime"`         // 操作时间
}

type LinkedAdminOperLogAdminDept struct {
	gmeta.Meta `orm:"table:admin_dept"`
	DeptId     int64  `orm:"dept_id" json:"deptId"`     // 部门id
	DeptName   string `orm:"dept_name" json:"deptName"` // 部门名称
}

type AdminOperLogListRes struct {
	OperId         uint64                       `json:"operId"`
	Title          string                       `json:"title"`
	RequestMethod  string                       `json:"requestMethod"`
	OperName       string                       `json:"operName"`
	DeptName       string                       `json:"deptName"`
	LinkedDeptName *LinkedAdminOperLogAdminDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperUrl        string                       `json:"operUrl"`
	OperIp         string                       `json:"operIp"`
	OperLocation   string                       `json:"operLocation"`
	OperParam      string                       `json:"operParam"`
	OperTime       *gtime.Time                  `json:"operTime"`
}
