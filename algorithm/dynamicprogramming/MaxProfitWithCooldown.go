package dynamicprogramming

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 题目：309. 最佳买卖股票时机含冷冻期
// 难度：中等
// 描述：给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 思路：动态规划
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
// 解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。
// 参考：https://labuladong.gitbook.io/algo/mu-lu-ye-2/mu-lu-ye-4/tuan-mie-gu-piao-wen-ti
func maxProfitWithCooldown(prices []int) int {
	dp0, dp1 := 0, math.MinInt32
	dp_prev_0 := 0 // 代表dp[i-2][0]
	for i := 0; i < len(prices); i++ {
		tmp := dp0
		dp0 = base.Max(dp0, dp1+prices[i])
		dp1 = base.Max(dp1, dp_prev_0-prices[i])
		dp_prev_0 = tmp
	}
	return dp0
}
