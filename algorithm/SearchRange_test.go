package algorithm

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := [][]interface{}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).([]int)
		if got := searchRange(p1, p2); !reflect.DeepEqual(got, want) {
			t.Errorf("searchRange(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
