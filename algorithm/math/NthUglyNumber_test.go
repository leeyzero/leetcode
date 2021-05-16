package math

import (
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	tests := [][]int{
		{1, 1},
		{2, 2},
		{10, 12},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		if got := nthUglyNumber(p1); got != want {
			t.Errorf("nthUglyNumber(%v).got:%v want:%v", p1, got, want)
		}
	}
}
