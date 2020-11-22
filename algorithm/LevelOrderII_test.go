package algorithm

import (
	"reflect"
	"testing"
)

func TestLevelOrderII(t *testing.T) {
	in := []int{3, 9, -1, -1, 20, 15, -1, -1, 7, -1, -1}
	root := makeTreeNode(in)
	want := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
	if got := levelOrderII(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderII(%v).got:%v want:%v", in, got, want)
	}
}