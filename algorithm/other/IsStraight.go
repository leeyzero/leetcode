package other

import (
	"sort"
)

// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
// 题目：剑指 Offer 61. 扑克牌中的顺子
// 难度：简单
// 描述：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
// 思路：
// 1. 除大小王，所有牌无重复
// 2. 牌中最大值 - 最小值 小于5 max - min < 5
func isStraight(nums []int) bool {
	if len(nums) != 5 {
		return false
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	joker := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[len(nums)-1]-nums[joker] < 5
}
