package algorithm

import (
	"testing"
)

func TestMinArray(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(int)
		if got := minArray(in); got != want {
			t.Errorf("minArray(%v).got:%v want:%v", in, got, want)
		}
	}
}