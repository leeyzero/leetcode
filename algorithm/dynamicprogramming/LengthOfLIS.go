package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 题目：300. 最长递增子序列
// 难度：中等
// 思路：动态规划
// dp[i]表示以nums[i]为结尾的最长递增子序列
// dp[i+1]可以通过找到所有比nums[i+1]小的数结尾的最长递增子序列加上nums[i+1]
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = base.Max(dp[i], dp[j]+1)
			}
		}
	}

	ans := 0
	for i := 0; i < len(dp); i++ {
		ans = base.Max(ans, dp[i])
	}
	return ans
}
