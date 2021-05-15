package search

import (
	"testing"
)

func TestMovingCount(t *testing.T) {
	tests := [][]int{
		{2, 3, 1, 3},
		{3, 1, 0, 1},
		{16, 8, 4, 15},
	}
	for _, test := range tests {
		p1 := test[0]
		p2 := test[1]
		p3 := test[2]
		want := test[3]
		got := movingCount(p1, p2, p3)
		if got != want {
			t.Errorf("movingCount(%v, %v, %v).got:%v want:%v", p1, p2, p3, got, want)
		}
	}
}
