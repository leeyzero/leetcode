package search

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := [][]interface{}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 2, -1}, {0, 1, -1}}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).([][]int)
		if got := threeSum(nums); !reflect.DeepEqual(got, want) {
			t.Errorf("threeSum(%v).got:%v want:%v", nums, got, want)
		}
	}
}
