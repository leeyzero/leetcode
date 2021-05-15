package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
	}
	for _, test := range tests {
		if got, want := add(test[0], test[1]), test[2]; got != want {
			t.Errorf("add(%v, %v).got:%v want:%v", test[0], test[1], got, want)
		}
	}
}
