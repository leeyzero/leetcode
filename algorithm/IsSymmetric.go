package algorithm

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
// 解题思路：对树中任意两个对称结点L和R，一定有：
// - L.Val == R.Val
// - L.Left.Val == R.Right.Val
// - L.Right.Val == R.Left.Val
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(lnode *TreeNode, rnode *TreeNode) bool {
	if lnode == nil && rnode == nil {
		return true
	}

	if lnode == nil || rnode == nil || lnode.Val != rnode.Val {
		return false
	}
	return isSymmetricCore(lnode.Left, rnode.Right) && isSymmetricCore(lnode.Right, rnode.Left)
}

// TODO: 使用迭代实现
