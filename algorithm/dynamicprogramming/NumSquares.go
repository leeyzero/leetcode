package dynamicprogramming

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/perfect-squares/
// 题目：279. 完全平方数
// 难度：中等
// 描述：给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// 思路：动态规划
// dp[i]表示组成正整数i的最少完全平方数
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = base.Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
