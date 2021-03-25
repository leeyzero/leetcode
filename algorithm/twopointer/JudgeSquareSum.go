package twopointer

import (
	"math"
)

// https://leetcode-cn.com/problems/sum-of-square-numbers/
// 题目：633. 平方数之和
// 难度：中等
// 思路：双指标
func judgeSquareSum(c int) bool {
	low, high := 0, (int)(math.Sqrt(float64(c)))
	for low <= high {
		sum := low*low + high*high
		if sum == c {
			return true
		} else if sum < c {
			low++
		} else {
			high--
		}
	}
	return false
}
