package binarysearch

import (
	"testing"
)

func TestSearchInRotateSortedArray2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 1, 13, 1, 1, 1, 1, 1}, 13, true},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).(bool)
		if got := searchInRotateSortedArray2(nums, target); got != want {
			t.Errorf("searchInRotateSortedArray2(%v, %v).got:%v want:%v", nums, target, got, want)
		}
	}
}
