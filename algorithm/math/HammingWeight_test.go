package math

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := [][]int{
		{0, 0},
		{9, 2},
		{8, 1},
	}
	for _, test := range tests {
		in := (uint32)(test[0])
		if got := hammingWeight(in); got != test[1] {
			t.Errorf("hammingWeight(%v).got:%v want:%v", in, got, test[1])
		}
	}
}
