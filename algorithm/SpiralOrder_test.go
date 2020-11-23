package algorithm

import (
	"testing"
	"reflect"
)

func TestSpiralOrder(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{[]int{1}}, []int{1}},
		{[][]int{[]int{1, 2}}, []int{1, 2}},
		{[][]int{[]int{1}, []int{2}}, []int{1, 2}},
		{[][]int{[]int{1, 2}, []int{3, 4}}, []int{1, 2, 4, 3}},
		{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{[]int{6, 9, 7}}, []int{6, 9, 7}},
		{[][]int{[]int{1}, []int{2}, []int{3}}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		in := (test[0]).([][]int)
		want := (test[1]).([]int)
		if got := spiralOrder(in); !reflect.DeepEqual(got, want) {
			t.Errorf("spiralOrder(%v).got:%v want:%v", in, got, want)
		}
	}
}