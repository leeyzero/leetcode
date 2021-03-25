package twopointer

import (
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	tests := [][]interface{}{
		{5, true},
		{3, false},
		{4, true},
		{2, true},
		{1, true},
	}
	for _, test := range tests {
		c := (test[0]).(int)
		want := (test[1]).(bool)
		if got := judgeSquareSum(c); got != want {
			t.Errorf("judgeSquareSum(%v).got:%v want:%v", c, got, want)
		}
	}
}
