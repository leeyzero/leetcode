package twopointer

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}
	for _, test := range tests {
		height := (test[0]).([]int)
		want := (test[1]).(int)
		if got := maxArea(height); got != want {
			t.Errorf("maxArea(%v).got:%v want:%v", height, got, want)
		}
	}
}
