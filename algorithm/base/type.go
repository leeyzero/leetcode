package base

// 单向链表结点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
