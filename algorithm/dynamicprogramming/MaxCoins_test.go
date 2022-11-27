package dynamicprogramming

import "testing"

func TestMaxCoins(t *testing.T) {
	cases := [][]interface{}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
	}
	for _, c := range cases {
		p1 := c[0].([]int)
		want := c[1].(int)
		if got := maxCoins(p1); got != want {
			t.Errorf("maxCoins(%v) failed.got:%v want:%v", p1, got, want)
		}
	}
}
