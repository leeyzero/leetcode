package greedy

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).([]int)
		want := (test[2]).(int)
		if got := findContentChildren(p1, p2); got != want {
			t.Errorf("findContentChildren(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
