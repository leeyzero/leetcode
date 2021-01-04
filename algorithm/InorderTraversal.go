package algorithm

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalIterative(root)
}

func inorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, top.Val)
			curr = top.Right
		}
	}
	return ans
}

func inorderTraversalRecursive(root *TreeNode) []int {
	ans := []int{}
	inorderTraversalRecursiveCore(root, &ans)
	return ans
}

func inorderTraversalRecursiveCore(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	inorderTraversalRecursiveCore(node.Left, ans)
	*ans = append(*ans, node.Val)
	inorderTraversalRecursiveCore(node.Right, ans)
}
