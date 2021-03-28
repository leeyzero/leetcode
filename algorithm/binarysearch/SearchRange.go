package binarysearch

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 题目：34. 在排序数组中查找元素的第一个和最后一个位置
// 难度：中等
// 思路：二分法
func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}

	lBound := binarySearchLeftBound(nums, target)
	rBound := binarySearchRightBound(nums, target)
	return []int{lBound, rBound}
}

// target的左边界
func binarySearchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// target的右边界
func binarySearchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
