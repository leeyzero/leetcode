package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// 题目：剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// 难度：简单
// 思路：递归
func lowestBinarySearchTreeCommonAncestor(root *base.TreeNode, p *base.TreeNode, q *base.TreeNode) *base.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if (root.Val > p.Val) && (root.Val > q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	} else if (root.Val < p.Val) && (root.Val < q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
