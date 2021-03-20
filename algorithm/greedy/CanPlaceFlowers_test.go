package greedy

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).(bool)
		if got := canPlaceFlowers(p1, p2); got != want {
			t.Errorf("canPlaceFlowers(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
