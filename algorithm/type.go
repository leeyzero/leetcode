package algorithm

// 单向链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 多叉树节点定义
type Node struct {
	Val int
	Children []*Node
}