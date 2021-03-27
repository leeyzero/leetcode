package sort

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 1, 1, 2, 2, 3, 4}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		k := (test[1]).(int)
		want := (test[2]).([]int)
		if got := topKFrequent(nums, k); !reflect.DeepEqual(got, want) {
			t.Errorf("topKFrequent(%v, %v).got:%v, want:%v", nums, k, got, want)
		}
	}
}
