package algorithm

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 12},
	}
	for _, test := range tests {
		p1 := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := maxValue(p1); got != want {
			t.Errorf("maxValue(%v).got:%v want:%v", p1, got, want)
		}
	}
}
