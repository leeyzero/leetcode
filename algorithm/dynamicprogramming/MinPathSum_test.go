package dynamicprogramming

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
	}
	for _, test := range tests {
		grid := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := minPathSum(grid); got != want {
			t.Errorf("minPathSum(%v).got:%v want:%v", grid, got, want)
		}
	}
}
