/*
* @desc:用户在线状态
* @company:
* @Author:
* @Date:   2023/1/10 15:08
 */

package model

// AdminUserOnlineParams 用户在线状态写入参数
type AdminUserOnlineParams struct {
	UserAgent string
	Uuid      string
	Token     string
	Username  string
	Ip        string
}
