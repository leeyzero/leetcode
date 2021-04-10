package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// 二叉树的后续遍历
func postorderTraversal(root *base.TreeNode) []int {
	return postorderTraversalIterative(root)
}

// https://www.geeksforgeeks.org/iterative-postorder-traversal-using-stack/
// https://en.wikipedia.org/wiki/Tree_traversal
func postorderTraversalIterative(root *base.TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*base.TreeNode{}
	curr, lastVisitNode := root, (*base.TreeNode)(nil)
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			top := stack[len(stack)-1]
			// if right child exist and traversing node from left, then move to right child
			if top.Right != nil && top.Right != lastVisitNode {
				curr = top.Right
			} else {
				ans = append(ans, top.Val)
				stack = stack[:len(stack)-1]
				lastVisitNode = top
			}
		}
	}
	return ans
}

func postorderTraversalRecursive(root *base.TreeNode) []int {
	ans := []int{}
	postorderTraversalRecursiveCore(root, &ans)
	return ans
}

func postorderTraversalRecursiveCore(node *base.TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	postorderTraversalRecursiveCore(node.Left, ans)
	postorderTraversalRecursiveCore(node.Right, ans)
	*ans = append(*ans, node.Val)
}
