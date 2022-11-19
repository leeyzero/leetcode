package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestFlatten(t *testing.T) {
	in := []int{1, 2, 5, 3, 4, -1, 6}
	root, err := base.UnmarshalTreeNodeByLevelorder(in)
	if err != nil {
		t.Fatalf("base.UnmarshalTreeNodeByLevelorder(%v) failed due to %v", in, err)
	}

	flatten(root)

	got := base.PreOrderTraverseTreeNode(root)
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("flatten failed.got:%v want:%v", got, want)
	}
}
