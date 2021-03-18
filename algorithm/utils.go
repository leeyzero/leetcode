package algorithm

func makeLinkList(vals []int) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return sentinel.Next
}

func forward(head *ListNode, length int) *ListNode {
	curr := head
	for i := 0; curr != nil && i < length; i++ {
		curr = curr.Next
	}
	return curr
}

// 按前序遍历反序列化
// [1, 2, 4, -1, -1, 5, -1, -1, 3, -1, -1]
//      1
//   2    3
// 4  5
func makeTreeNode(vals []int) *TreeNode {
	root := (*TreeNode)(nil)
	makeTreeNodeAux(&root, &vals)
	return root
}

func makeTreeNodeAux(node **TreeNode, vals *[]int) {
	if len(*vals) <= 0 {
		return
	}

	val := (*vals)[0]
	*vals = (*vals)[1:]
	if val == -1 {
		return
	}

	*node = &TreeNode{val, nil, nil}
	makeTreeNodeAux(&(*node).Left, vals)
	makeTreeNodeAux(&(*node).Right, vals)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func convertToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func preorderTreeNode(root *TreeNode) []int {
	ans := []int{}
	preorderTreeNodeAux(root, &ans)
	return ans
}

func preorderTreeNodeAux(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	*ans = append(*ans, node.Val)
	preorderTreeNodeAux(node.Left, ans)
	preorderTreeNodeAux(node.Right, ans)
}
