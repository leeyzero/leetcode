package search

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// 题目：111. 二叉树的最小深度
// 难度：简单
// 描述：给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
// 说明：叶子节点是指没有子节点的节点。
// 思路：DFS或BFS
// 方法一：DFS
func minDepth(root *base.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MaxInt32
	minDepthCore(root, 0, &ans)
	return ans
}

func minDepthCore(node *base.TreeNode, curr int, ans *int) {
	if node == nil {
		return
	}

	curr++
	if node.Left == nil && node.Right == nil {
		*ans = base.Min(*ans, curr)
		return
	}

	minDepthCore(node.Left, curr, ans)
	minDepthCore(node.Right, curr, ans)
}

// 方法二：BFS
func minDepth2(root *base.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	q := []*base.TreeNode{root}
	for len(q) > 0 {
		var nextLevel []*base.TreeNode
		for i := 0; i < len(q); i++ {
			if q[i].Left == nil && q[i].Right == nil {
				return depth
			}
			if q[i].Left != nil {
				nextLevel = append(nextLevel, q[i].Left)
			}
			if q[i].Right != nil {
				nextLevel = append(nextLevel, q[i].Right)
			}
		}
		q = nextLevel
		depth++
	}
	return depth
}
