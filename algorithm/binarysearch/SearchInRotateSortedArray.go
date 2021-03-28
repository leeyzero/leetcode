package binarysearch

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 题目：33. 搜索旋转排序数组
// 难度：中等
// 思路：二分法的变种，根据左右区间是否有序收缩搜索区间
func searchInRotateSortedArray(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}

		// 左区间有序
		if nums[lo] < nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] { // 右区间有序
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			break
		}
	}
	return -1
}
