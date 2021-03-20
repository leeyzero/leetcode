package dynamicprogramming

import (
	"testing"
)

func TestMaxProfit2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(int)
		if got := maxProfit2(p1); got != want {
			t.Errorf("maxProfit2(%v).got:%v want:%v", p1, got, want)
		}
	}
}
