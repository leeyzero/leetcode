package binarysearch

import (
	"testing"
)

func TestFindMin2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := findMin2(nums); got != want {
			t.Errorf("findMin2(%v).got:%v want:%v", nums, got, want)
		}
	}
}
