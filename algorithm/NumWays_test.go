package algorithm

import (
	"testing"
)

func TestNumWays(t *testing.T) {
	tests := [][]int{
		{0, 1},
		{1, 1},
		{2, 2},
		{7, 21},
	}
	for _, test := range tests {
		if r := numWays(test[0]); r != test[1] {
			t.Errorf("test numWays(%d).got:%d want:%d", test[0], r, test[1])
		}
	}
}
