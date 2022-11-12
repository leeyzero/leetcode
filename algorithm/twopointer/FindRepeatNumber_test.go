package twopointer

import (
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	type Table struct {
		in  []int
		out int
	}
	table := []Table{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2},
		{[]int{3, 4, 2, 1, 1, 0}, 1},
	}
	for _, v := range table {
		if got, want := findRepeatNumber(v.in), v.out; got != want {
			t.Errorf("findRepeatNumber(%v) failed.got:%v want:%v", v.in, got, want)
		}
	}
}
