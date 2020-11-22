package algorithm

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// 解题思路：使用一个队列存储每一层节点，每次将一层的节点出队并将其子节点入队
func levelOrderII(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []*TreeNode{}
		curr := []int{}
		for _, node := range queue {
			curr = append(curr, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		ans = append(ans, curr)
		queue = level
	}
	return ans
}