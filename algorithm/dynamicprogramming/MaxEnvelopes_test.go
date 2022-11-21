package dynamicprogramming

import (
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	cases := [][]interface{}{
		{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}
	for _, c := range cases {
		p1 := c[0].([][]int)
		want := c[1].(int)
		if got := maxEnvelopes(p1); got != want {
			t.Errorf("maxEnvelopes(%v) failed.got:%v want:%v", p1, got, want)
		}
	}
}
