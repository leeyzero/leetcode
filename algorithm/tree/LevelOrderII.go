package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// 题目：剑指 Offer 32 - II. 从上到下打印二叉树 II
// 难度：简单
// 思路：使用一个队列存储每一层节点，每次将一层的节点出队并将其子节点入队
func levelOrderII(root *base.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	queue := []*base.TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		level := []*base.TreeNode{}
		ans = append(ans, []int{})
		for _, node := range queue {
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		queue = level
	}
	return ans
}
