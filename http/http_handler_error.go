package http

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// NewErr 新增err
func NewErr(code int, msg string, remark ...string) error {
	r := ""
	if len(remark) > 0 {
		r = remark[0]
	}
	return gerror.NewCode(gcode.New(code, r, nil), msg)
}

// NewErrWithCode 新增err， 指定code
func NewErrWithCode(code gcode.Code, msg string, remark ...string) error {
	if len(remark) == 0 {
		return gerror.NewCode(code, msg)
	}
	return gerror.NewCode(gcode.New(code.Code(), remark[0], nil), msg)

}

// NotAuthorizedErr 鉴权失败err
func NotAuthorizedErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeAuthorizedErr, msg)
}

// NotFoundErr 无此资源err
func NotFoundErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeNotFoundErr, msg)
}

// InvalidParamErr 入参错误err
func InvalidParamErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeInvalidParamErr, msg)
}

// InternalErr 服务端错误err
func InternalErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeInternalErr, msg)
}
