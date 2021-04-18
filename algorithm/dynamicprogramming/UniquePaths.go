package dynamicprogramming

// https://leetcode-cn.com/problems/unique-paths/
// 题目：62. 不同路径
// 难度：中等
// 思路：动态规划（二维）
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for i, j := 0, 0; j < n; j++ {
		dp[i][j] = 1
	}
	for i, j := 0, 0; i < m; i++ {
		dp[i][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
