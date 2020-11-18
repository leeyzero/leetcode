package algorithm

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}
