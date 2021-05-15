package twopointer

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
// 题目：剑指 Offer 57. 和为s的两个数字
// 难度：简单
// 思路：双指针
func twoSum(nums []int, target int) []int {
	ans := []int{}
	if len(nums) <= 1 {
		return ans
	}

	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if target == sum {
			ans = []int{nums[i], nums[j]}
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return ans
}
