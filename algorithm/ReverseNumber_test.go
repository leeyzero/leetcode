package algorithm

import (
	"testing"
)

func TestReverseNumber(t *testing.T) {
	tests := [][]int{
		// {123, 321},
		{-123, -321},
		// {120, 21},
	}
	for _, test := range tests {
		p1 := test[0]
		if got, want := reverseNumber(p1), test[1]; got != want {
			t.Errorf("reverse(%v).got:%v want:%v", p1, got, want)
		}
	}
}