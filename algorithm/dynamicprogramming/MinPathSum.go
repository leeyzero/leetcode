package dynamicprogramming

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/minimum-path-sum/
// 题目：64. 最小路径和
// 难度：中等
// 描述：给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + base.Min(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[m-1][n-1]
}
