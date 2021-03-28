package binarysearch

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
// 题目：154. 寻找旋转排序数组中的最小值 II
// 难度：困难
// 思路：二分法，该题是153的一个变种，区别是数组中可能有重复元素
// 旋转排序数组nums可以被拆分为2个排序数组nums1,nums2，并且nums1任一元素 >= nums2任一元素；
// 因此，考虑二分法寻找此两数组的分界点nums[i] (即第2个数组的首个元素)
func findMin2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] { // 说明mid一定是在第一个排序数组中，i一定满足mid < i <= hi
			lo = mid + 1
		} else if nums[mid] < nums[hi] { // 说明mid一定是在第二个排序数组中，i一定满足lo < i <= mid
			hi = mid
		} else {
			hi = hi - 1
		}
	}
	return nums[lo]
}
