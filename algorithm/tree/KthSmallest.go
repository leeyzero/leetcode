package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 题目：230. 二叉搜索树中第K小的元素
// 难度：中等
// 思路：二仅搜索树中序遍历元素递增
func kthSmallest(root *base.TreeNode, k int) int {
	var ans int
	kthSmallestCore(root, &k, &ans)
	return ans
}

func kthSmallestCore(node *base.TreeNode, k *int, ans *int) {
	if node == nil || *k <= 0 {
		return
	}

	kthSmallestCore(node.Left, k, ans)
	*k--
	if *k == 0 {
		*ans = node.Val
		return
	}
	kthSmallestCore(node.Right, k, ans)
}
