package dynamicprogramming

import (
	"testing"
)

func TestMinInsertions(t *testing.T) {
	cases := [][]interface{}{
		{"zzazz", 0},
		{"mbadm", 2},
		{"leetcode", 5},
	}
	for _, c := range cases {
		p1 := c[0].(string)
		want := c[1].(int)
		if got := minInsertions(p1); got != want {
			t.Errorf("minInsertions(%v) failed.got:%v want:%v", p1, got, want)
		}
	}
}
