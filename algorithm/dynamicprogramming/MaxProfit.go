package dynamicprogramming

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
// 题目：剑指 Offer 63. 股票的最大利润
// 难度：中等
// 思路：动态规划
// dp[i]表示第i天的最大收益
// dp[i] = max(dp[i-1], prices[i] - min_cost(0...i))
// 初始态：dp[0] = 0
func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	cost := math.MaxInt32
	profit := 0
	for i := 1; i < len(prices); i++ {
		cost = base.Min(cost, prices[i-1])
		profit = base.Max(profit, prices[i]-cost)
	}
	return profit
}
