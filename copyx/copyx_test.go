package copyx

import (
	"testing"
)

func TestCopyx(t *testing.T) {

	type SA struct {
		A string
	}

	type SB struct {
		A string
	}

	t.Run("*struct to *struct", func(t *testing.T) {
		var a *SA
		var b *SB
		sb := Copy(a, b)

		t.Log("result sb: ", sb)
	})

	t.Run("*struct to *struct", func(t *testing.T) {
		a := &SA{A: "xxxxxxxxxxx"}
		sb := Copy(a, new(SB))

		t.Log("result sb: ", sb)
	})

	t.Run("*struct to *struct", func(t *testing.T) {
		a := &SA{A: "xxxxxxxxxxx"}
		b := &SB{}
		sb := Copy(a, b)

		t.Log("result sb: ", sb)
	})
}
