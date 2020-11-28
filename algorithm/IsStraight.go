package algorithm

import (
	"sort"
)
// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
// 解题思路：
// 1. 除大小王，所有牌无重复
// 2. 牌中最大值 - 最小值 小于5 max - min < 5
func isStraight(nums []int) bool {
    if len(nums) != 5 {
        return false
    }

    sort.Slice(nums, func (i, j int) bool {
        return nums[i] < nums[j]
    })

    joker := 0
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
    return nums[len(nums)-1] - nums[joker] < 5
}