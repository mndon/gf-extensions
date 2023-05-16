package httpx

import (
	"database/sql"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	handleResponse(r)
}

func handleResponse(r *ghttp.Request) {
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg    string
		remark string
		ctx    = r.Context()
		err    = r.GetError()
		res    = r.GetHandlerResponse()
		code   = gerror.Code(err)
	)

	if err != nil {
		// 捕获框架内置code， 映射成自定义code
		switch code {
		case gcode.CodeNil: //未知报错
			switch err {
			case sql.ErrNoRows:
				code = CodeNotFoundErr
			default:
				code = CodeInternalErr
			}
			remark = code.Message()
		case gcode.CodeValidationFailed: // 参数校验失败
			code = CodeInvalidParamErr
			remark = err.Error()
		default:
			remark = code.Message()
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		// 捕获http错误响应码， 映射成自定义错误类型
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusUnauthorized:
			code = CodeAuthorizedErr
		case http.StatusNotFound:
			code = CodeNotFoundErr
		case http.StatusForbidden:
			code = CodeAuthorizedErr
		default:
			code = CodeInternalErr
		}
	} else {
		code = CodeOk
	}

	internalErr := r.Response.WriteJson(HandlerResponse{
		Status: code.Code(),
		Remark: remark,
		Msg:    msg,
		Data:   res,
	})
	if internalErr != nil {
		g.Log().Errorf(ctx, `%+v`, internalErr)
	}
}
