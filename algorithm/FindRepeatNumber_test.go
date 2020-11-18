package algorithm

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
	}
	for _, v := range table {
		if findRepeatNumber(v.in) != v.out {
			t.Errorf("test fail.case:%+v", v)
		}
	}
}
