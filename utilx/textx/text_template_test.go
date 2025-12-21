package textx

import (
	"testing"
)

func TestTplRender(t *testing.T) {
	tpl := "<div>1{a}2</div>"
	vars := map[string]any{
		"a": "b",
	}
	result := TplRender(tpl, vars)
	t.Log(result)
}
