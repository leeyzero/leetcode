package dynamicprogramming

import "testing"

func TestMaxProfitWithCooldown(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 0, 2}, 3},
	}
	for _, test := range tests {
		prices := (test[0]).([]int)
		want := (test[1]).(int)
		if got := maxProfitWithCooldown(prices); got != want {
			t.Errorf("maxProfitWithCooldown(%v).got:%v want:%v", prices, got, want)
		}
	}
}
