package http

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeOk              = gcode.New(2000, "", nil)
	CodeAuthorizedErr   = gcode.New(4010, AuthenticationFailedMsg.GetMsg(), nil)
	CodeNotFoundErr     = gcode.New(4040, NoSuchResourceMsg.GetMsg(), nil)
	CodeInvalidParamErr = gcode.New(4000, InvalidParamMsg.GetMsg(), nil)
	CodeInternalErr     = gcode.New(5000, InternalErrorMsg.GetMsg(), nil)

	LoginLimitErr = gcode.New(4011, LoginLimitMsg.GetMsg(), nil)
)
