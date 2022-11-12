package twopointer

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// 题目：剑指 Offer 03. 数组中重复的数字
// 描述：在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1
// 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// 思路：所有数字都是在0~n-1范围内，如果没有重复数据，则每个数字都能找到自己的位置
// 遍历数组，对每个数字放到属于它的位置上，如果这个位置已经已经放入了，说明该数字重复了
func findRepeatNumber(nums []int) int {
	for j := 0; j < len(nums); {
		if j == nums[j] {
			// nums[j]已经在正确的位置
			j++
			continue
		}

		n := nums[j]
		if nums[n] == n {
			// n所处的位置已经放置好了，说明重复
			return n
		}

		// 将n放置到正确的位置
		nums[n], nums[j] = nums[j], nums[n]
	}
	return -1
}
