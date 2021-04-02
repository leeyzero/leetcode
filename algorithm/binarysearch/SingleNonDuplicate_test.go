package binarysearch

import (
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := singleNonDuplicate(nums); got != want {
			t.Errorf("singleNonDuplicate(%v).got:%v want:%v", nums, got, want)
		}
	}
}
