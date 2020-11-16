package NumWays

import (
	"testing"
)

func TestNumWays(t *testing.T) {
	tests := [][]int{
		[]int{0, 1},
		[]int{1, 1},
		[]int{2, 2},
		[]int{7, 21},
	}
	for _, test := range tests {
		if r := numWays(test[0]); r != test[1] {
			t.Errorf("test numWays(%d).got:%d want:%d", test[0], r, test[1])
		}
	}
}