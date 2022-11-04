package http

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	CodeOk               = gcode.New(2000, "", "")
	CodeNotAuthorizedErr = gcode.New(4010, "authenticationFailed", "authenticationFailedRemark")
	CodeNotFoundErr      = gcode.New(4040, "noSuchResource", "noSuchResourceRemark")
	CodeInvalidParamErr  = gcode.New(4000, "invalidParam", "invalidParamRemark")
	CodeUnknownErr       = gcode.New(5000, "internalError", "internalErrorRemark")
)
