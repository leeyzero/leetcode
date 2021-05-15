package array

// https://leetcode-cn.com/problems/valid-mountain-array/
// 题目：941. 有效的山脉数组
// 难度：简单
// 描述：
// 给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。
// 让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
// arr.length >= 3
// 在 0 < i < arr.length - 1 条件下，存在 i 使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 思路：
// 我们从数组的最左侧开始向右扫描，直到找到第一个不满足arr[i] < arr[i+1] 的下标i，那么i就是这个数组的最高点的下标。
// 如果i = 0或者不存在这样的i（即整个数组都是单调递增的），那么就返回false。否则从i开始继续向右扫描，判断接下来的的下标j是否都
// 满足 arr[j]>arr[j+1]，若都满足就返回 true，否则返回false
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	var i, j int
	for i = 0; i < len(A)-1; i++ {
		if A[i+1] <= A[i] {
			break
		}
	}
	for j = len(A) - 1; j > 0; j-- {
		if A[j-1] <= A[j] {
			break
		}
	}
	return (i == j) && (i > 0) && (i < len(A)-1)
}
