package search

// https://leetcode-cn.com/problems/subsets/
// 题目：78. 子集
// 难度：中等
// 描述：
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 思路：回溯法
func subsets(nums []int) [][]int {
	ans := [][]int{}
	for count := 0; count <= len(nums); count++ {
		var curr []int
		backtrackingSubsets(nums, count, 0, curr, &ans)
	}
	return ans
}

func backtrackingSubsets(nums []int, count int, start int, curr []int, ans *[][]int) {
	if len(curr) == count {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrackingSubsets(nums, count, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
