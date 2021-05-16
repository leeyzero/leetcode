package math

import (
	"testing"
)

func TestCuttingRopeII(t *testing.T) {
	tests := [][]int{
		{2, 1},
		{10, 36},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		if got := cuttingRopeII(p1); got != want {
			t.Errorf("cuttingRopeII(%v).got:%v want:%v", p1, got, want)
		}
	}
}
