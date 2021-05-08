package tree

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 题目：124. 二叉树中的最大路径和
// 描述：路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其最大路径和 。
// 思路：
// 本质上是一棵二叉树的后序遍历，需要明确两个问题
// 1. 如何计算一个节点的最大值
// 2. 如何计算树中的最大路径
func maxPathSum(root *base.TreeNode) int {
	maxSum := math.MinInt32
	maxPathSumAux(root, &maxSum)
	return maxSum
}

func maxPathSumAux(node *base.TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// 分别计算左右节点的最大贡献值，只有在最大节点贡献值大于0时才选取
	left := base.Max(0, maxPathSumAux(node.Left, maxSum))
	right := base.Max(0, maxPathSumAux(node.Right, maxSum))

	// 节点的最大路径和取决于该节点的值加左右节点的最大贡献值
	*maxSum = base.Max(*maxSum, left+right+node.Val)

	// 返回该节点的最大贡献值
	return base.Max(left, right) + node.Val
}
