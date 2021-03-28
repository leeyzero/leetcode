package binarysearch

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
// 题目：81. 搜索旋转排序数组 II
// 难度：中等
// 思路：二分法，也是leetcode-33的变种，区别是数组中有重复元素
func searchInRotateSortedArray2(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}

		// 判断不出是前半部分有序还是后半部分有序：如果10111，11101
		if nums[mid] == nums[lo] {
			lo++
			continue
		}

		// 前半部分有序
		if nums[lo] < nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				// target在[lo, mid)范围内
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] <= nums[hi] { // 后半部分有序
			if target > nums[mid] && target <= nums[hi] { // target在(mid, hi]范围内
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			break
		}
	}
	return false
}
