package binarysearch

// https://leetcode-cn.com/problems/search-insert-position/
// 题目：35. 搜索插入位置
// 难度：简单
// 思路：二分法
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
