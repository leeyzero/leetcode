package twopointer

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 26. 删除有序数组中的重复项
// 难度：简单
// 描述：
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 思路：双指针
// 快慢指针分别指向数组起始位置，快指针每次移动一个单位，如果快指针指向的值和慢指针指向的值相等，说明元素重复，
// 移动快指针；如果不等，说明元素不再重复，将慢指针移动一个单元，并赋值为快指针的值，一直到数组结束
//
// 示例：
// [1, 1, 1] => [1]
// [1, 2, 2] => [1, 2]
// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] => [0, 1, 2, 3, 4]
func removeDuplicates(nums []int) int {
	lo := -1 // lo跟踪数组中最长不重复的下标，初始值为-1
	for hi := 0; hi < len(nums); hi++ {
		if lo < 0 || nums[lo] != nums[hi] {
			lo++
			nums[lo] = nums[hi]
		}
	}
	return lo + 1
}
