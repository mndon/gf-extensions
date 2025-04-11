package textx

import (
	"github.com/flosch/pongo2/v6"
)

// TplRender
// @Description: 文本模板渲染，同django/jinjia2 语法
// @param text
// @param vars
// @return string
func TplRender(text string, vars map[string]any) string {
	tpl, err := pongo2.FromString(text)
	if err != nil {
		return text
	}
	output, err := tpl.Execute(vars) // 输出渲染结果
	if err != nil {
		return text
	}
	return output
}
