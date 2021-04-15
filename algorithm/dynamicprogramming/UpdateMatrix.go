package dynamicprogramming

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/01-matrix/
// 题目：542. 01 矩阵
// 难度：中等
// 描述：给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
// 两个相邻元素间的距离为 1 。
// 思路：
// 左上到右下进行一次动态搜索，再从右下到左上进行一次动态搜索。两次动态搜索即可完成四个方向上的查找
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) <= 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] = base.Min(dp[i-1][j]+1, dp[i][j])
				}
				if j > 0 {
					dp[i][j] = base.Min(dp[i][j-1]+1, dp[i][j])
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				if i < m-1 {
					dp[i][j] = base.Min(dp[i+1][j]+1, dp[i][j])
				}
				if j < n-1 {
					dp[i][j] = base.Min(dp[i][j+1]+1, dp[i][j])
				}
			}
		}
	}
	return dp
}
