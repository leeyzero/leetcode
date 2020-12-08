package algorithm

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 解题思路：双指针
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
