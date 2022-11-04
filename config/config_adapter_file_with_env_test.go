package config

import (
	"context"
	"fmt"
	"github.com/mndon/gf-extensions/config/utils"
	"os"
	"testing"
)

var (
	ctx = context.TODO()
)

func TestName(t *testing.T) {
	t.Run("timestamp", func(t *testing.T) {
		os.Setenv("AA", "1")

		fmt.Println(utils.ExpandEnv("a${AA:-xx}b"))
		fmt.Println(utils.ExpandEnv("a${A:-xx}b"))
		fmt.Println(utils.ExpandEnv("a${A}b"))
	})
}
