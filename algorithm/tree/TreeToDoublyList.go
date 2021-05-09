package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// 题目：剑指 Offer 36. 二叉搜索树与双向链表
// 描述：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
// 难度：中等
// 思路：先找到树的最小值（最左结点）和最大值（最右结点），二叉搜索树的中序遍历是排好序的，可以按中序遍历构建双向链表
func treeToDoublyList(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return nil
	}

	head := getMinTreeNode(root)
	tail := getMaxTreeNode(root)
	pre := (*base.TreeNode)(nil)
	treeToDoublyListCore(root, &pre)
	head.Left, tail.Right = tail, head
	return head
}

func treeToDoublyListCore(node *base.TreeNode, pre **base.TreeNode) {
	if node == nil {
		return
	}

	treeToDoublyListCore(node.Left, pre)
	node.Left = *pre
	if *pre != nil {
		(*pre).Right = node
	}
	*pre = node
	treeToDoublyListCore(node.Right, pre)
}

func getMinTreeNode(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return nil
	}

	curr := root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func getMaxTreeNode(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return nil
	}

	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}
