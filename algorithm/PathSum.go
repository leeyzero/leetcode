package algorithm

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
// 解题思路：先序遍历，用一个栈记录路径，当到达叶子节点，并且路径的和为指定值时，保存路径
// 注意golang中slice的内部实现原理，保存路径时需要做下深拷贝
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