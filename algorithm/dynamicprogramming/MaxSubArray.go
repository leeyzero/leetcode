package algorithm

import (
	"math"
)

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// 题目：剑指 Offer 42. 连续子数组的最大和
// 难度：简单
// 解题思路：f(i)表示元素nums[i]为结尾的连续数组最大和
// f(i) = nums[i]， i == 0 或则 f(i-1) <= 0
// f(i) = f(i-1) + nums[i]， i > 0 并且 f(i-1) > 0
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	ans := math.MinInt32
	for i, _ := range nums {
		if i == 0 || dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// 由于f(i)只跟f(i-1)有关系，只需要用两个变量记录f(i)和最大连续数组和即可
func maxSubArray2(nums []int) int {
	ans := math.MinInt32
	curr := 0
	for i, _ := range nums {
		if i == 0 || curr <= 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}

		if curr > ans {
			ans = curr
		}
	}
	return ans
}
