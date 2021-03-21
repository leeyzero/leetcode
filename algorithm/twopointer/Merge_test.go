package twopointer

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		p3 := (test[2]).([]int)
		p4 := (test[3]).(int)
		want := (test[4]).([]int)
		merge(p1, p2, p3, p4)
		if !reflect.DeepEqual(p1, want) {
			t.Errorf("merge(%v, %v, %v, %v).got:%v want:%v", p1, p2, p3, p4, p1, want)
		}
	}
}
