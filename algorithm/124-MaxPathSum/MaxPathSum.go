package MaxSumPath

import (
    "math"
)

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 本质上是一棵二叉树的后序遍历，需要明确两个问题
// 1. 如何计算一个节点的最大值
// 2. 如何计算树中的最大路径
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    maxPathSumAux(root, &maxSum)
    return maxSum
}

func maxPathSumAux(node *TreeNode, maxSum *int) int {
    if node == nil {
        return 0
    }

    // 分别计算左右节点的最大贡献值，只有在最大节点贡献值大于0时才选取
    left := max(0, maxPathSumAux(node.Left, maxSum))
    right := max(0, maxPathSumAux(node.Right, maxSum))

    // 节点的最大路径和取决于该节点的值加左右节点的最大贡献值
    *maxSum = max(*maxSum, left + right + node.Val)

    // 返回该节点的最大贡献值
    return max(left, right) + node.Val
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
