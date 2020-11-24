package algorithm

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 2, 2, 2, 5, 4, 2}, 2},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(int)
		if got := majorityElement(in); got != want {
			t.Errorf("majorityElement(%v).got:%v want:%v", in, got, want)
		}
	}
}