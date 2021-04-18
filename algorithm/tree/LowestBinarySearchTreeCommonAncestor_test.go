package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestlowestBinarySearchTreeCommonAncestor(t *testing.T) {
	p1 := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	root, err := base.UnmarshalTreeNodeByLevelorder(p1)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", p1, err)
	}

	p, q := root.Left, root.Right
	ancestor := lowestBinarySearchTreeCommonAncestor(root, p, q)
	if ancestor == nil {
		t.Fatalf("lowestBinarySearchTreeCommonAncestor(%v, %v, %v) got nil", root.Val, p.Val, q.Val)
	}

	want := root
	if ancestor != want {
		t.Fatalf("lowestBinarySearchTreeCommonAncestor(%v, %v, %v). got:%v want:%v",
			root.Val, p.Val, q.Val, ancestor.Val, root.Val)
	}
}
