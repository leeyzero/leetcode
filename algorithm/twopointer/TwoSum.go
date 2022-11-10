package twopointer

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
// 题目：剑指 Offer 57. 求有序数组中和为s的两个数字
// 难度：简单
// 思路：双指针
func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if target == sum {
			return []int{nums[i], nums[j]}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
