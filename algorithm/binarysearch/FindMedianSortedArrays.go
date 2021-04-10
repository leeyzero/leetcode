package binarysearch

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// 题目：4. 寻找两个正序数组的中位数
// 难度：困难
// 思路：
// 方法一：合并两个有序数组，找出中位数，时间复杂度是O(m+n)，空间复杂度是O(m+n)
// 方法二：维护两个指针，初始化为两个数组下标0的位置，每次将指向较小值的指针向后移动一位（如果一个指针已到达末尾，
// 移动另一个指针即可），直到到达中位数的位置。时间复杂度O(m+n)，空间复杂度O(1)
// 方法三：
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0.0
}
