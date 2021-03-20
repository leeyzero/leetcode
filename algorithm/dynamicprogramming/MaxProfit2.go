package dynamicprogramming

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 题目：122. 买卖股票的最佳时机 II
// 难度：简单
// 思路：动态规则
// 定义状态:
//  - dp[i][0]表示第i天交易完后手里没有股票的最大利润;
//  - dp[i][1]表示第i天交易完后手里持有一支股票的最大利润;
//  - i = [0, len(prices))
// 状态转移：
//  - dp[i][0] = max{dp[i−1][0], dp[i−1][1] + prices[i]}
//  - dp[i][1] = max{dp[i−1][1], dp[i−1][0] − prices[i]}
// 初始状态：
//  - dp[0][0] = 0
//  - dp[0][1] = -prices[0]
// 由于当前状态只跟前一个状态有关，所以可以使用两个变更来表示，很显然，对于最后一天j, dp[j][0]肯定是要大于dp[j][1]的
func maxProfit2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		nextdp0 := base.Max(dp0, dp1+prices[i])
		nextdp1 := base.Max(dp1, dp0-prices[i])
		dp0, dp1 = nextdp0, nextdp1
	}
	return dp0
}
