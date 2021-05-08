package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMirrorTree(t *testing.T) {
	in := []int{4, 2, 1, -1, -1, 3, -1, -1, 7, 6, -1, -1, 9, -1, -1}
	want := []int{4, 7, 9, 6, 2, 3, 1}
	root := base.MakeTreeNode(in)
	if got := base.PreOrderTraverseTreeNode(mirrorTree(root)); !reflect.DeepEqual(got, want) {
		t.Errorf("mirrorTree(%v).got:%v want:%v", in, got, want)
	}
}
