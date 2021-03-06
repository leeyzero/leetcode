package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		got := make([]int, len(p1))
		copy(got, p1)
		want := (test[1]).([]int)
		MergeSort(got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("insertSort(%v).got:%v want:%v", p1, got, want)
		}
	}
}
