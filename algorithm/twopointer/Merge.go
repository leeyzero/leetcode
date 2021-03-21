package twopointer

// https://leetcode-cn.com/problems/merge-sorted-array/
// 题目：88. 合并两个有序数组
// 难度：简单
// 解题思路：双指标，从尾部开始比较
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m+n {
		return
	}

	i, j, k := m-1, n-1, len(nums1)-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[i], nums1[k] = nums1[k], nums1[i]
			i--
		} else {
			nums2[j], nums1[k] = nums1[k], nums2[j]
			j--
		}
	}
	for ; i >= 0; i, k = i-1, k-1 {
		nums1[i], nums1[k] = nums1[k], nums1[i]
	}
	for ; j >= 0; j, k = j-1, k-1 {
		nums2[j], nums1[k] = nums1[k], nums2[j]
	}
}
