package search

import "sort"

// https://leetcode-cn.com/problems/3sum/
// 题目：15. 三数之和
// 难度：中等
// 描述：
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 思路：排序 + 回溯
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n int, start int, target int) [][]int {
	ans := [][]int{}
	if n < 2 || len(nums) < 2 {
		return ans
	}

	if n == 2 {
		lo, hi := start, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				ans = append(ans, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				ans = append(ans, arr)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ans
}
