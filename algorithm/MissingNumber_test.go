package algorithm

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := [][]interface{}{
		{[]int{0, 1, 3}, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9}, 8},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(int)
		if got := missingNumber(p1); got != want {
			t.Errorf("missingNumber(%v).got:%v want:%v", p1, got, want)
		}
	}
}