package binarysearch

// https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
// 540. 有序数组中的单一元素
// 中等
// 思路：二分法，根据mid所处下标判断
func singleNonDuplicate(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
