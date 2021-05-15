package twopointer

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 26. 删除有序数组中的重复项
// 难度：简单
// 描述：
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 思路：双指针
func removeDuplicates(nums []int) int {
	j := -1
	for i := 0; i < len(nums); i++ {
		if j < 0 || nums[j] != nums[i] {
			j += 1
			nums[j] = nums[i]
		}
	}
	return j + 1
}
