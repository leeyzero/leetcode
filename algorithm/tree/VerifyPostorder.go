package tree

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
// 题目：剑指 Offer 33. 二叉搜索树的后序遍历序列
// 难度：中等
// 描述：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
// 思路：
// 后续遍历先遍历左子树，再遍历右子树，最后才遍历根节点。由于时二叉搜索树，左子树的所有节点均小于
// 跟节点，右子树的所有节点均大于右子树，输入数组的根是数组的最后一个元素。
// 可以利用这个特性判断序列是否时二叉搜索树的顺序
func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	root := postorder[len(postorder)-1]
	rbegin := 0
	// 从[0,len-1)中找到第一个大于root的位置rbegin,
	// 则[0,rbegin)为root的左子树, [rbegin, len-1)为root的右子树
	for ; rbegin < len(postorder)-1; rbegin++ {
		if postorder[rbegin] > root {
			break
		}
	}
	// 判断右子树所有节点是否大于root
	for i := rbegin; i < len(postorder)-1; i++ {
		if postorder[i] < root {
			return false
		}
	}

	// 递归判断左右子树
	return verifyPostorder(postorder[0:rbegin]) && verifyPostorder(postorder[rbegin:len(postorder)-1])
}
