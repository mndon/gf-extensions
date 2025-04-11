package textx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTplRender(t *testing.T) {
	tpl := "1{{a}}2"
	vars := map[string]any{
		"a": "a",
	}
	result := TplRender(tpl, vars)
	t.Log(result)
	assert.Equal(t, "1a2", result)
}
