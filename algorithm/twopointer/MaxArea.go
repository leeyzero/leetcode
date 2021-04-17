package twopointer

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/container-with-most-water/
// 题目：11. 盛最多水的容器
// 难度：中等
// 描述：
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器。
// 思路：双指针
func maxArea(height []int) int {
	ans := 0
	lo, hi := 0, len(height)-1
	for lo < hi {
		ans = base.Max(ans, base.Min(height[hi], height[lo])*(hi-lo))
		if height[lo] <= height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return ans
}
