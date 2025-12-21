package textx

import (
	"github.com/slongfield/pyfmt"
)

// TplRender
// @Description: 文本模板渲染，同django/jinjia2 语法
// @param text
// @param vars
// @return string
func TplRender(text string, vars map[string]any) string {
	output, err := pyfmt.Fmt(text, vars)
	if err != nil {
		return text
	}
	return output
}
