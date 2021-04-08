package search

import (
	"testing"
)

func TestShortestBridge(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{0, 1}, {1, 0}}, 1},
		{[][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}, 2},
		{[][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}, 1},
	}
	for _, test := range tests {
		A := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := shortestBridge(A); got != want {
			t.Errorf("shortestBridge(%v).got:%v want:%v", A, got, want)
		}
	}
}
