package twopointer

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 题目：剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 难度：简单
// 思路：双指针，时间复杂度O(N), 空间复杂度O(1)
func exchange(nums []int) []int {
	j := -1 // 用j指向奇数的最后一个元素，初始为-1，表示没有奇数
	for i, num := range nums {
		if num%2 == 1 {
			// 奇数
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func exchange2(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 1 {
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}
	return nums
}
