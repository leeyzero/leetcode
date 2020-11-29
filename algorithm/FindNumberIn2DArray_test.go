package algorithm

import (
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1,   4,  7, 11, 15},
  		{2,   5,  8, 12, 19},
  		{3,   6,  9, 16, 22},
  		{10, 13, 14, 17, 24},
  		{18, 21, 23, 26, 30},
	}
	tests := [][]interface{}{
		{matrix, 5, true},
		{matrix, 20, false},
	}
	for _, test := range tests {
		p1 := (test[0]).([][]int)
		p2 := (test[1]).(int)
		want := (test[2]).(bool)
		if got := findNumberIn2DArray(p1, p2); got != want {
			t.Errorf("findNumberIn2DArray(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}