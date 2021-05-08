package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
// 题目：剑指 Offer 27. 二叉树的镜像
// 描述：请完成一个函数，输入一个二叉树，该函数输出它的镜像。
// 思路：本质上是先序遍历，交换左右节点，递归处理左右子树
func mirrorTree(root *base.TreeNode) *base.TreeNode {
	mirrorTreeCore(root)
	return root
}

func mirrorTreeCore(node *base.TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	mirrorTree(node.Left)
	mirrorTree(node.Right)
}
