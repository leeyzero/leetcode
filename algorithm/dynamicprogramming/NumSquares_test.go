package dynamicprogramming

import "testing"

func TestNumSquares(t *testing.T) {
	tests := [][]int{
		{12, 3},
		{13, 2},
	}
	for _, test := range tests {
		n := test[0]
		want := test[1]
		if got := numSquares(n); got != want {
			t.Errorf("numSquares(%v).got:%v want:%v", n, got, want)
		}
	}
}
