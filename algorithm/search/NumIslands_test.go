package search

import "testing"

func TestNumIslands(t *testing.T) {
	tests := [][]interface{}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1,
		},
	}

	for _, test := range tests {
		grid := (test[0]).([][]byte)
		want := (test[1]).(int)
		if got := numIslands(grid); got != want {
			t.Errorf("numIslands(%v).got:%v want:%v", grid, got, want)
		}
	}
}
