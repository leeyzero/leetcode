package dynamicprogramming

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := [][]int{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, test := range tests {
		m := test[0]
		n := test[1]
		want := test[2]
		if got := uniquePaths(m, n); got != want {
			t.Errorf("uniquePaths(%v, %v).got:%v want:%v", m, n, got, want)
		}
	}
}
