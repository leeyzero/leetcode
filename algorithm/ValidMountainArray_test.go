package algorithm

import (
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	type Table struct {
		in  []int
		out bool
	}
	table := []Table{
		{[]int{0, 3, 2, 1}, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	}
	for _, v := range table {
		if validMountainArray(v.in) != v.out {
			t.Errorf("test case fail.case:%+v", v)
		}
	}
}
