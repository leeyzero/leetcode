package greedy

// https://leetcode-cn.com/problems/non-decreasing-array/
// 题目：665. 非递减数列
// 难度：简单
// 思路：贪心
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return count <= 1
}
