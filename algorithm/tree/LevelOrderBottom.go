package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 题目：107. 二叉树的层序遍历 II
// 难度：中等
// 描述：给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
func levelOrderBottom(root *base.TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*base.TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ans = append(ans, []int{})
		count := len(queue)
		for k := 0; k < count; k++ {
			node := queue[0]
			queue = queue[1:]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
