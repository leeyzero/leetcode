package algorithm

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 解题思路：使用patiton思想，遍历数组，将奇数放前面，偶数放后面
func exchange(nums []int) []int {
	j := -1 // 用j指向奇数的最后一个元素，初始为-1，表示没有奇数
	for i, num := range nums {
		if num % 2 == 1 {
			// 奇数
			j++
			nums[i], nums[j] = nums[j], nums[i]
		} 
	}
	return nums
}