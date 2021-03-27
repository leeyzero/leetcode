package sort

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 215. 数组中的第K个最大元素
// 难度：中等
// 思路：双指针，二分法
func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	lo, hi, target := 0, len(nums)-1, len(nums)-k
	for lo < hi {
		mid := quickSelect(nums, lo, hi)
		if mid == target {
			return nums[mid]
		} else if mid < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0
}

// 快速排序中的partition原理
func quickSelect(arr []int, lo int, hi int) int {
	i := lo
	for j := lo + 1; j <= hi; j++ {
		if arr[j] <= arr[lo] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[lo], arr[i] = arr[i], arr[lo]
	return i
}
