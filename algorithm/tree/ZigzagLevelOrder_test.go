package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestZigzagLevelOrder(t *testing.T) {
	in := []int{3, 9, -1, -1, 20, 15, -1, -1, 7}
	want := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	root := base.MarshalTreeNodeByPreorder(in)
	if got := zigzagLevelOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("zigzagLevelOrder(%v).got:%v want:%v", in, got, want)
	}
}
