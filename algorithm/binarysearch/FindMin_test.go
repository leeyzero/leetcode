package binarysearch

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{1}, 1},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := findMin(nums); got != want {
			t.Errorf("findMin(%v).got:%v want:%v", nums, got, want)
		}
	}
}
