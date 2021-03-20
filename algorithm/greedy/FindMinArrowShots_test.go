package greedy

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}, 2},
		{[][]int{{1, 9}, {7, 16}, {2, 5}, {7, 12}, {9, 11}, {2, 10}, {9, 16}, {3, 9}, {1, 3}}, 2},
		{[][]int{{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13}}, 3},
	}
	for _, test := range tests {
		points := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := findMinArrowShots(points); got != want {
			t.Errorf("findMinArrowShots(%v).got:%v want:%v", points, got, want)
		}
	}
}
