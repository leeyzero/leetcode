package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestLowestCommonAncestor(t *testing.T) {
	p1 := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root, err := base.UnmarshalTreeNodeByLevelorder(p1)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", p1, err)
	}

	p := root.Left
	q := root.Left.Right.Right
	commonAncestor := lowestCommonAncestor(root, p, q)
	if commonAncestor == nil {
		t.Fatalf("lowestCommonAncestor(%v, %v, %v).got nil", root.Val, p.Val, q.Val)
	}

	want := p
	if commonAncestor != want {
		t.Errorf("lowestCommonAncestor(%v, %v, %v).got:%v want:%v",
			root.Val, p.Val, q.Val, commonAncestor.Val, want.Val)
	}
}
