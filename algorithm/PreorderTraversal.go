package algorithm

// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalIterative(root)
}

// 迭代
func preorderTraversalIterative(node *TreeNode) []int {
	ans := []int{}
	if node == nil {
		return ans
	}

	stack := []*TreeNode{}
	curr := node
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			ans = append(ans, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			// 左子树为空
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curr = top.Right
		}
	} 
	return ans
}

// 递归
func preorderTraversalRecursive(root *TreeNode) []int {
	ans := []int{}
	preorderTraversalRecursiveCore(root, &ans)
	return ans
}

func preorderTraversalRecursiveCore(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	*ans = append(*ans, node.Val)
	preorderTraversalRecursiveCore(node.Left, ans)
	preorderTraversalRecursiveCore(node.Right, ans)
}