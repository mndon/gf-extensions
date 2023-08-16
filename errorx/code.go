package errorx

import "github.com/gogf/gf/v2/errors/gcode"

var (
	AuthenticationFailedMsg = I18nMsg{"验证用户身份失败，请重新登录", "User authentication failed, please log in again"}
	NoSuchResourceMsg       = I18nMsg{"无此内容，请咨询客服", "No such content, please consult customer service"}
	InvalidParamMsg         = I18nMsg{"请正确输入参数", "Please enter the parameters correctly"}
	InternalErrorMsg        = I18nMsg{"系统错误[错误代码:5000]，多次失败，请咨询客服", "System error [error code: 5000], multiple failures, please consult customer service"}
	LoginLimitMsg           = I18nMsg{"登录异常，请检查是否有多个客户端同时登录", "Login of multiple devices, please check and log in again"}
)

// gcode.code => res.status
// gcode.message => res.remark
// error.Error() => res.msg
var (
	CodeOk            = gcode.New(2000, "", nil)
	CodeAuthorizedErr = gcode.New(4010, AuthenticationFailedMsg.GetMsg(), nil)
	CodeNotFoundErr   = gcode.New(4040, NoSuchResourceMsg.GetMsg(), nil)
	CodeBadRequestErr = gcode.New(4000, InvalidParamMsg.GetMsg(), nil)
	CodeInternalErr   = gcode.New(5000, InternalErrorMsg.GetMsg(), nil)
	CodeLoginLimitErr = gcode.New(4011, LoginLimitMsg.GetMsg(), nil)
)
