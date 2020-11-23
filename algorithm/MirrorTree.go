package algorithm

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
// 解题思路：交换左右节点，递归处理左右子树
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
