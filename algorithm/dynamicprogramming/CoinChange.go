package dynamicprogramming

import (
	"math"

	"github.com/leeyzero/leetcode/algorithm/base"
)

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

	// 数组大小初始化为amount+1, 初始值也为amount+1
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 外层循环遍历所有状态的所有取值
	for i := 1; i <= amount; i++ {
		// 内层循环求所有选择的最小值
		for _, coin := range coins {
			// 币值大于要兑换的零钱，无解，跳过
			if coin > i {
				continue
			}
			dp[i] = base.Min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 使用递归解法，自上而下
func coinChange2(coins []int, amount int) int {
	var closure func(n int) int
	closure = func(n int) int {
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}

		ans := math.MaxInt
		for _, coin := range coins {
			subproblem := closure(n - coin)
			if subproblem < 0 {
				// 子问题无解，跳过
				continue
			}
			ans = base.Min(ans, subproblem+1)
		}
		if ans == math.MaxInt {
			ans = -1
		}
		return ans
	}
	return closure(amount)
}

// 带备忘录的递归
func coinChange3(coins []int, amount int) int {
	memo := map[int]int{}
	var closure func(n int) int
	closure = func(n int) int {
		// 先查备忘录
		if v, ok := memo[n]; ok {
			return v
		}

		// base case
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}

		// 在子问题中选择最优解
		ans := math.MaxInt
		for _, coin := range coins {
			subproblem := closure(n - coin)
			if subproblem < 0 {
				// 子问题无解，跳过
				continue
			}
			ans = base.Min(ans, subproblem+1)
		}
		if ans == math.MaxInt {
			ans = -1
		}

		// 写入备忘录
		memo[n] = ans
		return ans
	}

	return closure(amount)
}
