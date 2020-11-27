package algorithm

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
// 解题思路：二分法
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left) >> 1
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