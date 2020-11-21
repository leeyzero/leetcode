package algorithm

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum(root *TreeNode, sum int) [][]int {
	path := []int{}
	ans := [][]int{}
	pathSumAux(root, sum, 0, path, &ans)
	return ans
}

func pathSumAux(node *TreeNode, sum int, curr int, path []int, ans *[][]int) {
	if node == nil {
		return
	}
	
	
	path = append(path, node.Val)
	curr += node.Val
	if node.Left == nil && node.Right == nil && curr == sum {
		cpy := make([]int, len(path))
		copy(cpy, path)
		*ans = append(*ans, cpy)
	}

	pathSumAux(node.Left, sum, curr, path, ans)
	pathSumAux(node.Right, sum, curr, path, ans)
	path = path[:len(path)-1]
}