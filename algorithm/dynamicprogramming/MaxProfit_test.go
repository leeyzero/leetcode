package dynamicprogramming

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := [][]interface{}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(int)
		if got := maxProfit(in); got != want {
			t.Errorf("maxProfit(%v).got:%v want:%v", in, got, want)
		}
	}
}
