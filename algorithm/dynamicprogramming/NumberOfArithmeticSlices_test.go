package dynamicprogramming

import (
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := numberOfArithmeticSlices(nums); got != want {
			t.Errorf("numberOfArithmeticSlices(%v).got:%v want:%v", nums, got, want)
		}
	}
}
