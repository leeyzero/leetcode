package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
// 题目：剑指 Offer 55 - I. 二叉树的深度
// 难度：简单
// 思路：本质上是后序遍历
func maxDepth(root *base.TreeNode) int {
	if root == nil {
		return 0
	}

	ldepth := maxDepth(root.Left)
	rdepth := maxDepth(root.Right)
	return 1 + base.Max(ldepth, rdepth)
}
