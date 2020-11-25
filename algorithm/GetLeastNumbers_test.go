package algorithm

import (
	"reflect"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 2, 1}, 2, []int{1, 2}},
		{[]int{0, 1, 2, 1}, 1, []int{0}},
		{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8, []int{0, 0, 1, 1, 2, 2, 2, 3}},
		{[]int{0, 0, 0, 2, 0, 5}, 0, []int{}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).([]int)
		if got := getLeastNumbers(p1, p2); !reflect.DeepEqual(got, want) {
			t.Errorf("getLeastNumbers(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
