package dynamicprogramming

import "testing"

func TestMaxProfit4(t *testing.T) {
	tests := [][]interface{}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
		{2, []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	}
	for _, test := range tests {
		k := (test[0]).(int)
		prices := (test[1]).([]int)
		want := (test[2]).(int)
		if got := maxProfit4(k, prices); got != want {
			t.Errorf("maxProfit4(%v, %v).got:%v want:%v", k, prices, got, want)
		}
	}
}
