package algorithm

import (
	"testing"
)

func TestLastRemaining(t *testing.T) {
	tests := [][]int{
		{5, 3, 3},
		{10, 17, 2},
	}
	for _, test := range tests {
		if got, want := lastRemaining(test[0], test[1]), test[2]; got != want {
			t.Errorf("lastRemaining(%v, %v).got:%v want:%v", test[1], test[1], got, want)
		}
	}
}
