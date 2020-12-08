package algorithm

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
// 解题思路：动态规划
// f(i, j)表示从(0,0)到(i,j)的最大价值
// 初始化：f(0, 0) = grid[0][0]
// 状态转移：
// f(i, j) = grid[i][j] + f(i, j-1) (i = 0, j > 0)
// f(i, j) = grid[i][j] + f(i-1, j) (i > 0, j = 0)
// f(i, j) = grid[i][j] + max(f(i-1, j), f(i, j-1) (x > 0, j > 0)
func maxValue(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[n-1][m-1]
}

