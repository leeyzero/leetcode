package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// 二叉树的前序遍历
func preorderTraversal(root *base.TreeNode) []int {
	return preorderTraversalIterative(root)
}

// https://www.geeksforgeeks.org/iterative-preorder-traversal/
// 1. create a stack, and push root node to stack
// 2. do following while stack is empty
//   a) pop an item from stack and visit it
//   b) push right child of popped item to stack
//   c) push left child of popped item to stack
// Note: Right child is pushed before left child to make sure that left subtree is processed first.
func preorderTraversalIterative(node *base.TreeNode) []int {
	ans := []int{}
	if node == nil {
		return ans
	}

	stack := []*base.TreeNode{node}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return ans
}

// 递归
func preorderTraversalRecursive(root *base.TreeNode) []int {
	ans := []int{}
	preorderTraversalRecursiveCore(root, &ans)
	return ans
}

func preorderTraversalRecursiveCore(node *base.TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	*ans = append(*ans, node.Val)
	preorderTraversalRecursiveCore(node.Left, ans)
	preorderTraversalRecursiveCore(node.Right, ans)
}
