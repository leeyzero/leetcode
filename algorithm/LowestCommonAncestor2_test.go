package algorithm

import (
	"testing"
)

func TestLowestCommonAncestor2(t *testing.T) {
	in := []int{3, 5, 6, -1, -1, 2, 7, -1, -1, 4, -1, -1, 1, 0, -1, -1, 8, -1, -1}
	root := makeTreeNode(in)
	tests := [][]interface{}{
		{in, root.Left, root.Right, root},
		{in, root.Left, root.Left.Right.Right, root.Left},
	}
	for _, test := range tests {
		p := (test[1]).(*TreeNode)
		q := (test[2]).(*TreeNode)
		want := (test[3]).(*TreeNode)
		if got := lowestCommonAncestor2(root, p, q); got != want {
			t.Errorf("lowestCommonAncestorII(%v, %v, %v).got:%v want:%v", in, p.Val, q.Val, got, want)
		}
	}
}