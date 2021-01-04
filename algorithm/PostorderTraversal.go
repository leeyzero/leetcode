package algorithm

// 二叉树的后续遍历
func postorderTraversal(root *TreeNode) []int {
	return postorderTraversalIterative(root)
}

func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Right != nil && len(stack) > 0 && curr.Right == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			stack = append(stack, curr)
			curr = curr.Right
		} else {
			ans = append(ans, curr.Val)
			curr = nil
		}
	}
	return ans
}

func postorderTraversalRecursive(root *TreeNode) []int {
	ans := []int{}
	postorderTraversalRecursiveCore(root, &ans)
	return ans
}

func postorderTraversalRecursiveCore(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	postorderTraversalRecursiveCore(node.Left, ans)
	postorderTraversalRecursiveCore(node.Right, ans)
	*ans = append(*ans, node.Val)
}