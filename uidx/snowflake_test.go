package uidx

import (
	"fmt"
	"testing"
)

func TestUid(t *testing.T) {

	t.Run("idgen", func(t *testing.T) {
		Init(0)
		fmt.Println(NextId())
	})
}
