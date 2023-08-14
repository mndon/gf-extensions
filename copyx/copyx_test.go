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
		sb := SB{}
		err := Copy(&sa, &sb)
		if err != nil {
			t.Error(err)
		}
		t.Log("result sb: ", sb)
	})

	t.Run("*struct to *struct, ignore IgnoreEmpty", func(t *testing.T) {
		sb := SB{
			A: "BBBBB",
			B: 1,
			C: &Item{},
			D: []int{1},
		}
		err := Copy(&sa, &sb)
		if err != nil {
			t.Error(err)
		}
		t.Log("result sb: ", sb)
	})

	t.Run("*struct to **struct", func(t *testing.T) {
		var sb *SB
		err := Copy(&sa, &sb)
		if err != nil {
			t.Error(err)
		}
		t.Log("result sb: ", sb)
	})

	t.Run("*struct to **struct, ignore IgnoreEmpty", func(t *testing.T) {
		sb := &SB{
			A: "BBBBB",
			B: 1,
			C: &Item{},
			D: []int{1},
		}
		err := Copy(&sa, &sb)
		if err != nil {
			t.Error(err)
		}
		t.Log("result sb: ", sb)
	})

	t.Run("*struct to *struct, filed differences", func(t *testing.T) {
		type SAA struct {
			A *string
			B *int
			C *Item
			D *[]int
		}

		saaA := "xxxxx"
		saaD := []int{1, 2}

		saa := SAA{
			A: &saaA,
			B: nil,
			C: nil,
			D: &saaD,
		}

		sb := &SB{
			A: "BBBBB",
			B: 1,
			C: &Item{},
			D: []int{1},
		}
		err := Copy(&saa, &sb)
		if err != nil {
			t.Error(err)
		}
		t.Log("result sb: ", sb)
	})
}
