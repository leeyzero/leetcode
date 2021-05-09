package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
// 题目：剑指 Offer 34. 二叉树中和为某一值的路径
// 难度：中等
// 描述：输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
// 思路：先序遍历，用一个栈记录路径，当到达叶子节点，并且路径的和为指定值时，保存路径
// 注意golang中slice的内部实现原理，保存路径时需要做下深拷贝
func pathSum(root *base.TreeNode, sum int) [][]int {
	path := []int{}
	ans := [][]int{}
	pathSumAux(root, sum, 0, path, &ans)
	return ans
}

func pathSumAux(node *base.TreeNode, sum int, curr int, path []int, ans *[][]int) {
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
}
