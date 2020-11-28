package algorithm

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 解题思路：二叉树中序遍历为递增序列，中序遍历倒序为递减序列
func kthLargest(root *TreeNode, k int) int {
	ans := (*TreeNode)(nil)
	kthLargestCore(root, &k, &ans)
	if ans != nil {
		return ans.Val
	}
	return 0
}

func kthLargestCore(node *TreeNode, k *int, res **TreeNode) {
	if node == nil || *k <= 0 {
		return
	}
	
	kthLargestCore(node.Right, k, res)
	(*k)--
	if (*k) == 0 {
		*res = node
	}
	kthLargestCore(node.Left, k, res)
}