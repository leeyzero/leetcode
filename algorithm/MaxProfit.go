package algorithm

import (
	"math"
)

// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
// 解题思路：动态规划
// dp[i]表示第i天的最大收益
// dp[i] = max(dp[i-1], prices[i] - min_cost(0...i))
// 初始态：dp[0] = 0
func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0] = 0
	cost := math.MaxInt32
	for i := 1; i < len(prices); i++ {
		cost = min(cost, prices[i-1])
		dp[i] = max(dp[i-1], prices[i] - cost)
	}
	return dp[len(dp)-1]
}