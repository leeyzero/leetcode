package greedy

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// 题目：135. 分发糖果
// 链接：https://leetcode-cn.com/problems/candy/
// 难度：困难
// 思路：贪心
func candy(ratings []int) int {
	nums := make([]int, len(ratings))
	for i := range nums {
		nums[i] = 1
	}

	// left -> right
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}
	// right -> left
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			nums[i-1] = base.Max(nums[i-1], nums[i]+1)
		}
	}

	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
