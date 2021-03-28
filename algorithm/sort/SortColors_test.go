package sort

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		cpyNums := make([]int, len(nums))
		copy(cpyNums, nums)
		want := (test[1]).([]int)
		sortColors(nums)
		if !reflect.DeepEqual(nums, want) {
			t.Errorf("sortColors(%v).got:%v want:%v", cpyNums, nums, want)
		}
	}
}

func TestSortColors2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		cpyNums := make([]int, len(nums))
		copy(cpyNums, nums)
		want := (test[1]).([]int)
		sortColors2(nums)
		if !reflect.DeepEqual(nums, want) {
			t.Errorf("sortColors(%v).got:%v want:%v", cpyNums, nums, want)
		}
	}
}
