package algorithm

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
// 解题思路：后续遍历，避免重复计算
func isBalanced(root *TreeNode) bool {
	var depth int
	return isBalancedAux(root, &depth)
}

func isBalancedAux(node *TreeNode, depth *int) bool {
	if node == nil {
		return true
	}

	var left, right int
	if isBalancedAux(node.Left, &left) && isBalancedAux(node.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			*depth = 1 + max(left, right)
			return true
		}
	}
	return false
}
