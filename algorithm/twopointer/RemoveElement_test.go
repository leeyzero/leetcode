package twopointer

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).(int)
		if got := removeElement(p1, p2); got != want {
			t.Errorf("removeElement(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
