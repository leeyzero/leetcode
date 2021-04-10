package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 题目：103. 二叉树的锯齿形层序遍历
// 难度：中等
// 描述：给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
func zigzagLevelOrder(root *base.TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*base.TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		curr := []int{}
		count := len(queue)
		for k := 0; k < count; k++ {
			node := queue[0]
			queue = queue[1:]
			curr = append(curr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if i%2 != 0 {
			for l, r := 0, len(curr)-1; l < r; l, r = l+1, r-1 {
				curr[l], curr[r] = curr[r], curr[l]
			}
		}
		ans = append(ans, curr)
	}
	return ans
}
