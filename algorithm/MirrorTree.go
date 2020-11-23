package algorithm

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func mirrorTree(root *TreeNode) *TreeNode {
	mirrorTreeCore(root)
	return root
}

func mirrorTreeCore(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	mirrorTree(node.Left)
	mirrorTree(node.Right)
}
