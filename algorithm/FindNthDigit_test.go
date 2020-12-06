package algorithm

import (
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	tests := [][]int{
		{0, 0},
		{3, 3},
		{11, 0},
		{1001, 7},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		if got := findNthDigit(p1); got != want {
			t.Errorf("findNthDigit(%v).got:%v want:%v", p1, got, want)
		}
	}
}