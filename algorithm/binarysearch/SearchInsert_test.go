package binarysearch

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 7, 9, 10}, 9, 3},
		{[]int{3}, 2, 0},
		{[]int{3}, 4, 1},
	}
	for _, test := range tests {
		arr := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).(int)
		if got := searchInsert(arr, target); got != want {
			t.Errorf("binarySearch(%v, %v).got:%v want:%v", arr, target, got, want)
		}
	}
}
