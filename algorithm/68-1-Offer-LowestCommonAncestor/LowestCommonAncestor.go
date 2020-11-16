package LowestCommonAncestor

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// Definition for a binary tree node.
type TreeNode struct {
	Val int;
	Left *TreeNode;
	Right *TreeNode
};

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if (root == nil || p == nil || q == nil) {
		return nil;
	}

	if ((root.Val > p.Val) && (root.Val > q.Val)) {
		return lowestCommonAncestor(root.Left, p, q);
	} else if ((root.Val < p.Val) && (root.Val < q.Val)) {
		return lowestCommonAncestor(root.Right, p, q);
	} else {
		return root;
	}
}