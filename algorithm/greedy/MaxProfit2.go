package greedy

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 题目：122. 买卖股票的最佳时机 II
// 难度：简单
// 思路：贪心
// 1.单独交易日：今日价格 p1, 明日价格 p2, 今日买入, 明日卖出收益为 p2-p1, 负值为亏损
// 2.连续上涨交易日：连续上涨股票价格为 p1,p2,p3,...,pn, 最大收益为 pn-p1,等价于每天买卖
// 3.连续下降交易日：不买收益最大，即不亏损
func maxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
