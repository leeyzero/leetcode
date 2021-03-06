package tree

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// 429. N 叉树的层序遍历
// 难度：中等
// 思路：BFS
type Node struct {
	Val      int
	Children []*Node
}

func levelOrderNTree(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*Node{root}
	for i := 0; len(queue) > 0; i++ {
		count := len(queue)
		ans = append(ans, []int{})
		for k := 0; k < count; k++ {
			node := queue[0]
			queue = queue[1:]
			ans[i] = append(ans[i], node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
	}
	return ans
}
