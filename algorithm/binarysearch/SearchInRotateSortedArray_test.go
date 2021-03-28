package binarysearch

import (
	"testing"
)

func TestSearchInRotateSortedArray(t *testing.T) {
	tests := [][]interface{}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).(int)
		if got := searchInRotateSortedArray(nums, target); got != want {
			t.Errorf("searchInRotateSortedArray(%v, %v).got:%v want:%v", nums, target, got, want)
		}
	}
}
