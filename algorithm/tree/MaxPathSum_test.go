package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMaxPathSum(t *testing.T) {
	root := &base.TreeNode{1, nil, nil}
	root.Left = &base.TreeNode{2, nil, nil}
	root.Right = &base.TreeNode{3, nil, nil}
	if got, want := maxPathSum(root), 6; got != want {
		t.Errorf("maxPathSum(%v).got:%v want:%v", base.PreOrderTraverseTreeNode(root), got, want)
	}
}
