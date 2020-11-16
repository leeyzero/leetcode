package TwoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 7, 11, 15}, 9, []int{2, 7}},
		{[]int{10,26,30,31,47,60}, 40, []int{10, 30}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).([]int)
		if got := twoSum(nums, target); !reflect.DeepEqual(got, want) {
			t.Errorf("twoSum(%v, %v).got:%v want:%v", nums, target, got, want)
		}
	}
}