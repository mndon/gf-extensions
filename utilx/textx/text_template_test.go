package textx

import (
	"testing"
)

func TestTplRender(t *testing.T) {
	tpl := "1{{a}}2"
	vars := map[string]any{
		"a": "a",
	}
	result := TplRender(tpl, vars)
	t.Log(result)
}
