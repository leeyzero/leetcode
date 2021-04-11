package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/house-robber/
// 题目：198. 打家劫舍
// 难度：中等
// 描述：
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 思路：
// 定义一个数组dp，dp[i]表示抢劫到第i个房子时，可以抢劫的最大数量。dp[i]有两种情况，如果我们抢劫第i个房子，那么抢劫的金额为dp[i-2] + nums[i]
// 如果我们不抢劫第i个房子，那么抢劫的金额为dp[i-1]，所以递推式为：
// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = base.Max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}
