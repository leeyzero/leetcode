package sort

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		k := (test[1]).(int)
		want := (test[2]).(int)
		if got := findKthLargest(nums, k); got != want {
			t.Errorf("findKthLargest(%v, %v).got:%v want:%v", nums, k, got, want)
		}
	}
}
