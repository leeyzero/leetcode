package dynamicprogramming

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/maximal-square/
// 题目：221. 最大正方形
// 难度：中等
// 思路：动态规划
// dp[i][j] = k^2, 充分条件是dp[i-1][j], dp[i][j-1], dp[i-1][j-1]的值必须都不小于(k-1)^2
// 否则(i, j)位置不可能够成一个边长为k的正方形，如果三个值中的最小值为k-1，则(i,j)一定可以构成一个边长
// 为k的正方形
// dp[i][j] = 0, (matrix[i][j] = 0)
// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]), (matrix[i][j] = 1)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = base.Min(dp[i-1][j-1], base.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
			ans = base.Max(ans, dp[i][j])
		}
	}
	return ans * ans
}
