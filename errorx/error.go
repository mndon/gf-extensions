package errorx

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ErrFunc func() error

type ErrFuncf func(data ...string) error

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
	return NewErrWithCode(CodeAuthorizedErr, msg, remark...)
}

// NotFoundErr 无此资源err
func NotFoundErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeNotFoundErr, msg, remark...)
}

// BadRequestErr 错误请求err
func BadRequestErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeBadRequestErr, msg, remark...)
}

// InternalErr 服务端错误err
func InternalErr(msg string, remark ...string) error {
	return NewErrWithCode(CodeInternalErr, msg, remark...)
}

func BuildErr(code int, msg string, remark ...string) ErrFunc {
	return func() error {
		return NewErr(code, msg, remark...)
	}
}

func BuildErrf(code int, msg string, remark string) ErrFuncf {
	return func(data ...string) error {
		return NewErr(code, msg, fmt.Sprintf(remark, data))
	}
}

func BuildErrWithCode(code gcode.Code, msg string, remark ...string) ErrFunc {
	return func() error {
		if len(remark) == 0 {
			return gerror.NewCode(code, msg)
		}
		return gerror.NewCode(gcode.New(code.Code(), remark[0], nil), msg)
	}
}

func BuildErrfWithCode(code gcode.Code, msg string, remark string) ErrFuncf {
	return func(data ...string) error {
		if remark == "" {
			return gerror.NewCode(code, msg)
		}
		return gerror.NewCode(gcode.New(code.Code(), fmt.Sprintf(remark, data), nil), msg)
	}
}
