package math

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 4, 3, 3}, 4},
		{[]int{9, 1, 7, 9, 7, 9, 7}, 1},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(int)
		if got := singleNumber(in); got != want {
			t.Errorf("singleNumber(%v).got:%v want:%v", in, got, want)
		}
	}
}
