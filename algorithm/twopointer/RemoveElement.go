package twopointer

// https://leetcode-cn.com/problems/remove-element/
// 题目：27. 移除元素
// 难度：简单
// 描述：给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
// 思路：双指针
func removeElement(nums []int, val int) int {
	j := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
