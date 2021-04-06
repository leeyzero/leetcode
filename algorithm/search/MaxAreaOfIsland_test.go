package search

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
		{[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}, 6},
	}
	for _, test := range tests {
		grid := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := maxAreaOfIsland(grid); got != want {
			t.Errorf("maxAreaOfIsland(%v).got:%v want:%v", grid, got, want)
		}
	}
}

func TestMaxAreaOfIsland2(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			4,
		},
	}
	for _, test := range tests {
		grid := (test[0]).([][]int)
		want := (test[1]).(int)
		if got := maxAreaOfIsland2(grid); got != want {
			t.Errorf("maxAreaOfIsland(%v).got:%v want:%v", grid, got, want)
		}
	}
}
