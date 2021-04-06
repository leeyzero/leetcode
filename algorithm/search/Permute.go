package search

// https://leetcode-cn.com/problems/permutations/
// 题目：46. 全排列
// 难度：中等
// 描述：给定一个没有重复数字的序列，返回其所有可能的全排列。
// 思路：回溯法
func permute(nums []int) [][]int {
	ans := [][]int{}
	backtrackingPermute(nums, 0, &ans)
	return ans
}

func backtrackingPermute(nums []int, pos int, ans *[][]int) {
	if pos >= len(nums)-1 {
		curr := make([]int, len(nums))
		copy(curr, nums)
		*ans = append(*ans, curr)
		return
	}

	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]
		backtrackingPermute(nums, pos+1, ans)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
}
