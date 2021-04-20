package dynamicprogramming

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := [][]int{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{45, 134903163},
	}
	for _, test := range tests {
		if r := fib(test[0]); r != test[1] {
			t.Errorf("test fib fail.in:%d got:%d expect:%d", test[0], r, test[1])
		}
	}
}
