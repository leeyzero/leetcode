package algorithm

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head := getMinTreeNode(root)
	tail := getMaxTreeNode(root)
	pre := (*TreeNode)(nil)
	treeToDoublyListCore(root, &pre)
	head.Left, tail.Right = tail, head
    return head
}

func treeToDoublyListCore(node *TreeNode, pre **TreeNode) {
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

func getMinTreeNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	curr := root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func getMaxTreeNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}