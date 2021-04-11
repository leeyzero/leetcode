package dynamicprogramming

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := [][]interface{}{
		{2, 2},
		{3, 3},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		want := (test[1]).(int)
		if got := climbStairs(n); got != want {
			t.Errorf("climbStairs(%v).got:%v want:%v", n, got, want)
		}
	}
}

func TestClimbStairs2(t *testing.T) {
	tests := [][]interface{}{
		{2, 2},
		{3, 3},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		want := (test[1]).(int)
		if got := climbStairs2(n); got != want {
			t.Errorf("climbStairs(%v).got:%v want:%v", n, got, want)
		}
	}
}
