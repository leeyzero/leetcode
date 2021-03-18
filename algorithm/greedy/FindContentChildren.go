package greedy

import (
	"sort"
)

// 题目：455.分发饼干
// 链接：https://leetcode-cn.com/problems/assign-cookies/
// 难度：简单
// 思路：贪心
func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	children, cookies := 0, 0
	for children < len(g) && cookies < len(s) {
		if g[children] <= s[cookies] {
			children++
		}
		cookies++
	}
	return children
}
