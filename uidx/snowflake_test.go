package uidx

import (
	"fmt"
	"testing"
)

func TestUid(t *testing.T) {

	t.Run("idgen", func(t *testing.T) {
		fmt.Println(NextId())
	})
}
