package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 7, 9, 10}, 9, 3},
		{[]int{1}, 2, -1},
		{[]int{2}, 2, 0},
	}
	for _, test := range tests {
		arr := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).(int)
		if got := binarySearch(arr, target); got != want {
			t.Errorf("binarySearch(%v, %v).got:%v want:%v", arr, target, got, want)
		}
	}
}
