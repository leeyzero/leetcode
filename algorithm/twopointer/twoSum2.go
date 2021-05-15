package twopointer

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
// 题目：167. 两数之和 II - 输入有序数组
// 难度：简单
// 思路：双指针
func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
