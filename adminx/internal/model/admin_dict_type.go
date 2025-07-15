/*
* @desc:字典类型
* @company:
* @Author:
* @Date:   2022/3/18 11:56
 */

package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminDictTypeInfoRes struct {
	DictId    uint64      `orm:"dict_id,primary"  json:"dictId"`    // 字典主键
	DictName  string      `orm:"dict_name"        json:"dictName"`  // 字典名称
	DictType  string      `orm:"dict_type,unique" json:"dictType"`  // 字典类型
	Status    uint        `orm:"status"           json:"status"`    // 状态（0正常 1停用）
	Remark    string      `orm:"remark"           json:"remark"`    // 备注
	CreatedAt *gtime.Time `orm:"created_at"       json:"createdAt"` // 创建日期
}
