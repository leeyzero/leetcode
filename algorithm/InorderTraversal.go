package algorithm

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalIterative(root)
}

// https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
// 1. create an empty stack
// 2. iniialize current node as root
// 3. push the current to S and set current = current.Left until current is nil
// 4. do following until current is nil and stack is not empty
//   a). pop the top item from the stack
//   b). print the popped item, set current = popped.Right
//   c). go to step 3
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
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, curr.Val)
			curr = curr.Right
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
