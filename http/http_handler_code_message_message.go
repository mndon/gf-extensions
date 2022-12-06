package http

var (
	AuthenticationFailedMsg = I18nMsg{"验证用户身份失败，请重新登录", "User authentication failed, please log in again"}
	NoSuchResourceMsg       = I18nMsg{"无此内容，请咨询客服", "No such content, please consult customer service"}
	InvalidParamMsg         = I18nMsg{"请正确输入参数", "Please enter the parameters correctly"}
	InternalErrorMsg        = I18nMsg{"系统错误[错误代码:5000]，多次失败，请咨询客服", "System error [error code: 5000], multiple failures, please consult customer service"}
)
