package twopointer

import (
	"reflect"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, test := range tests {
		numbers := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).([]int)
		if got := twoSum2(numbers, target); !reflect.DeepEqual(got, want) {
			t.Errorf("twoSum2(%v, %v).got:%v want:%v", numbers, target, got, want)
		}
	}
}
