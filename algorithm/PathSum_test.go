package algorithm

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	in := []int{5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, 5, -1, -1, 1, -1, -1}
	sum := 22
	root := makeTreeNode(in)
	want := [][]int{
		{5, 4, 11, 2},
		{5, 8, 4, 5},
	}
	if got := pathSum(root, sum); !reflect.DeepEqual(got, want) {
		t.Errorf("pathSum(%v, %v).got:%v want:%v", in, sum, got, want)
	}
}