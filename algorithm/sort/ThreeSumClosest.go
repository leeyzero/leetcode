package sort

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/3sum-closest/
// 题目：16. 最接近的三数之和
// 难度：中等
// 描述：给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
// 思路：排序+双指针
// 先排序，固定第一个数a，在a之后找用两个指针pb,pc指定数a之后的区间首尾，
// 如果a + b + c == target, 直接返回
// 如果a + b + c < target, pb++
// 如果a + b + c > target, pc--
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32

	fnUpdateClosest := func(curr int) {
		if math.Abs(float64(curr-target)) < math.Abs(float64(ans-target)) {
			ans = curr
		}
	}
	for i := 0; i < len(nums); i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			fnUpdateClosest(sum)
			if sum == target {
				return sum
			} else if sum < target {
				lo++
			} else if sum > target {
				hi--
			}
		}
	}
	return ans
}
