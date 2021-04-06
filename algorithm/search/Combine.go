package search

// https://leetcode-cn.com/problems/combinations/
// 题目：77. 组合
// 难度：中等
// 描述：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// 思路：回溯法
func combine(n int, k int) [][]int {
	if k <= 0 || k > n {
		return nil
	}

	ans := [][]int{}
	curr := []int{}
	backtrackingCombine(n, k, 1, curr, &ans)
	return ans
}

func backtrackingCombine(n int, k int, pos int, curr []int, ans *[][]int) {
	if len(curr) >= k {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}

	for i := pos; i <= n; i++ {
		curr = append(curr, i)
		backtrackingCombine(n, k, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
