package tree

import (
	"math"

	. "github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
// 题目：剑指 Offer 55 - II. 平衡二叉树
// 难度：简单
// 描述：输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
// 思路：后续遍历，避免重复计算
func isBalanced(root *TreeNode) bool {
	return isBalancedAux(root) != -1
}

func isBalancedAux(node *TreeNode) int {
	if node == nil {
		return 0
	}

	ldepth := isBalancedAux(node.Left)
	if ldepth == -1 {
		return -1
	}

	rdepth := isBalancedAux(node.Right)
	if rdepth == -1 {
		return -1
	}

	if int(math.Abs(float64(ldepth-rdepth))) > 1 {
		return -1
	}

	return 1 + Max(ldepth, rdepth)
}
