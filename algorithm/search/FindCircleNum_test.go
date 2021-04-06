package search

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	tests := [][]interface{}{
		{
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			2,
		},
		{
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			3,
		},
	}
	for _, test := range tests {
		isConnected := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := findCircleNum(isConnected); got != want {
			t.Errorf("findCircleNum(%v).got:%v want:%v", isConnected, got, want)
		}
	}
}
