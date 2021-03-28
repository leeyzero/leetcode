package binarysearch

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
// 题目：153. 寻找旋转排序数组中的最小值
// 难度：中等
// 思路：二分法（变种题）
func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		// [lo, hi]单调递增
		if nums[lo] <= nums[hi] {
			return nums[lo]
		}

		mid := lo + (hi-lo)/2
		// [lo, mid]单调递增, 说明最小值在右区间
		if nums[lo] <= nums[mid] {
			lo = mid + 1
		} else {
			// 注意，这里hi取值为mid，因为取mid-1可能会露掉最小值，min指向的可能就是最小值
			hi = mid
		}
	}
	return -1
}
