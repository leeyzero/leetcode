package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 剑指 Offer 32 - I. 从上到下打印二叉树
// 难度：中等
// 描述：从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
func levelOrder(root *base.TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	queue := []*base.TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		ans = append(ans, curr.Val)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return ans
}
