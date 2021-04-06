package search

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := [][]interface{}{
		{
			[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			[][]int{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
	}
	for _, test := range tests {
		heights := (test[0]).([][]int)
		want := (test[1]).([][]int)
		if got := pacificAtlantic(heights); !reflect.DeepEqual(got, want) {
			t.Errorf("pacificAtlantic(%v).got:%v want:%v", heights, got, want)
		}
	}
}
