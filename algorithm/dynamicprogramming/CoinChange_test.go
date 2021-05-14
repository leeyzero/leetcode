package dynamicprogramming

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 1},
	}
	for _, test := range tests {
		coins := (test[0]).([]int)
		amount := (test[1]).(int)
		want := (test[2]).(int)
		if got := coinChange(coins, amount); got != want {
			t.Errorf("coinChange(%v, %v).got:%v want:%v", coins, amount, got, want)
		}
	}
}
