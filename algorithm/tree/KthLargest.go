package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 题目：剑指 Offer 54. 二叉搜索树的第k大节点
// 难度：简单
// 思路：二叉树中序遍历为递增序列，中序遍历倒序为递减序列
func kthLargest(root *base.TreeNode, k int) int {
	var ans int
	kthLargestCore(root, &k, &ans)
	return ans
}

func kthLargestCore(node *base.TreeNode, k *int, res *int) {
	if node == nil || *k <= 0 {
		return
	}

	kthLargestCore(node.Right, k, res)
	(*k)--
	if (*k) == 0 {
		*res = node.Val
	}
	kthLargestCore(node.Left, k, res)
}
