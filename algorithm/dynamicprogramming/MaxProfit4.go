package dynamicprogramming

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
// 题目：188. 买卖股票的最佳时机 IV
// 难度：困难
// 描述：给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 思路：动态规划
// 状态定义：dp[i][k][0 or 1]表示第i天最多进行了k次交易，手上持有或未持有股票
// base case:
// dp[-1][k][0] = 0
// 解释：因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0 。
// dp[-1][k][1] = -infinity
// 解释：还没开始的时候，是不可能持有股票的，用负无穷表示这种不可能。
// dp[i][0][0] = 0
// 解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0 。
// dp[i][0][1] = -infinity
// 解释：不允许交易的情况下，是不可能持有股票的，用负无穷表示这种不可能。
// 状态转移：
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// 解释：今天我没有持有股票，有两种可能：要么是我昨天就没有持有，然后今天选择 rest，所以我今天还是没有持有；要么是我昨天持有股票，但是今天我 sell 了，所以我今天没有持有股票了
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
// 解释：今天我持有着股票，有两种可能：要么我昨天就持有着股票，然后今天选择 rest，所以我今天还持有着股票；要么我昨天本没有持有，但今天我选择 buy，所以今天我就持有股票了。
// labuladong对股票买卖问题归纳总结的很好，
// 具体参考：https://labuladong.gitbook.io/algo/mu-lu-ye-2/mu-lu-ye-4/tuan-mie-gu-piao-wen-ti
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if k > n/2 {
		// 相当于无限次交易
		return maxProfitInfinite(prices)
	}

	dp := make([][][2]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 < 0 {
				// base case
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = base.Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = base.Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func maxProfitInfinite(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
