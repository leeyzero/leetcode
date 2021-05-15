package binarysearch

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
// 题目：剑指 Offer 53 - II. 0～n-1中缺失的数字
// 难度：简单
// 思路：二分法
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 结束循环后left指向的位置有两种情况：
	// 1. nums[left] != left (left < len(nums))
	// 2. left == len(nums)
	return left
}
