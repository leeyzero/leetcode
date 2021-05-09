package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
// 题目：剑指 Offer 28. 对称的二叉树
// 描述：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
// 难度：简单
// 思路：
// 本质上是前序遍历，对树中任意两个对称结点L和R，一定有：
// - L.Val == R.Val
// - L.Left.Val == R.Right.Val
// - L.Right.Val == R.Left.Val
func isSymmetric(root *base.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(lnode *base.TreeNode, rnode *base.TreeNode) bool {
	if lnode == nil && rnode == nil {
		return true
	}

	if lnode == nil || rnode == nil || lnode.Val != rnode.Val {
		return false
	}
	return isSymmetricCore(lnode.Left, rnode.Right) && isSymmetricCore(lnode.Right, rnode.Left)
}
