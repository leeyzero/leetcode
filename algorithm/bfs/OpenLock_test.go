package bfs

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	cases := [][]interface{}{
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	}
	for _, c := range cases {
		p1 := c[0].([]string)
		p2 := c[1].(string)
		want := c[2].(int)
		if got := openLock(p1, p2); got != want {
			t.Errorf("openLock(%v, %v) failed.got:%v want:%v", p1, p2, got, want)
		}
	}
}
