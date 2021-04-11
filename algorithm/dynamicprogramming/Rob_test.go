package dynamicprogramming

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := rob(nums); got != want {
			t.Errorf("rob(%v).got:%v want:%v", nums, got, want)
		}
	}
}
