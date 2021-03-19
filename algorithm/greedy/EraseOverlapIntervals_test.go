package greedy

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, 2},
	}
	for _, test := range tests {
		p1 := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := eraseOverlapIntervals(p1); got != want {
			t.Errorf("eraseOverlapIntervals(%v).got:%v want:%v", p1, got, want)
		}
	}
}
