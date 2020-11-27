package algorithm

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := [][]interface{}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 6, 0},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).(int)
		if got := search(p1, p2); got != want {
			t.Errorf("search(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}