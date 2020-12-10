package algorithm

import (
	"testing"
)

func TestMinNumber(t *testing.T) {
	tests := [][]interface{}{
		{[]int{10, 2}, "102"},
		{[]int{3, 30, 34, 5, 9}, "3033459"},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(string)
		if got := minNumber(p1); got != want {
			t.Errorf("minNumber(%v).got:%v want:%v", p1, got, want)
		}
	}
}