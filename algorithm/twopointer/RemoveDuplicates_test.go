package twopointer

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want1 := (test[1]).(int)
		want2 := (test[2]).([]int)
		got1 := removeDuplicates(p1)
		if got1 != want1 {
			t.Fatalf("removeDuplicates(%v).got1:%v want1:%v", p1, got1, want1)
		}
		got2 := p1[:got1]
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("removeDuplicates(%v).got2:%v want2:%v", p1, got2, want2)
		}
	}
}
