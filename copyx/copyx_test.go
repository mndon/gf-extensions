package copyx

import (
	"testing"
)

func TestCopyx(t *testing.T) {
	type Item struct {
		sub string
	}

	type SA struct {
		A string
		B int
		C *Item
		D []int
	}

	type SB struct {
		A string
		B int
		C *Item
		D []int
	}

	sa := SA{
		A: "q",
		B: 0,
		C: nil,
		D: nil,
	}

	t.Run("*struct to *struct", func(t *testing.T) {
		sb := Copy(sa, new(SB))

		t.Log("result sb: ", sb)
	})
}
