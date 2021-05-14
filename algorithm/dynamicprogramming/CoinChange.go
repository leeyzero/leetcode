package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/coin-change/
// 题目：322. 零钱兑换
// 难度：中等
// 描述：给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 你可以认为每种硬币的数量是无限的。
// 思路：动态规划
// dp[i]表示兑换总金额为i时，最小使用的硬币个数
// 状态转移：dp[i] = min(dp[i], dp[i-coin]+1)
func coinChange(coins []int, amount int) int {
	if len(coins) <= 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 2
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = base.Min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+2 {
		return -1
	}
	return dp[amount]
}
