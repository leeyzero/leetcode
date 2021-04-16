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

// 方法二：空间压缩
// 使用一维数组，对于第i行，当计算第j列时，j-1的值已经计算好保存在d[j-1]中，
// i-1行第j列的值上次已经计算出来保存到了dp[j]中,所有状态转移方程为
// dp[j] = grid[i][j] (i=0 and j=0)
// dp[j] = grid[i][j] + dp[j-1] (i=0)
// dp[j] = grid[i][j] + dp[j] (j=0)
// dp[j] = grid[i][j] + min(dp[j], dp[j-1])
func minPathSum2(grid [][]int) int {
    if len(grid) <= 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    dp := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                dp[j] = grid[i][j]
            } else if i == 0 {
                dp[j] = dp[j-1] + grid[i][j]
            } else if j == 0 {
                dp[j] = dp[j] + grid[i][j]
            } else {
                dp[j] = grid[i][j] + base.Min(dp[j], dp[j-1])
            }
        }
    }
    return dp[n-1]
}
