package twopointer

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problemset/all/?search=340
// 题目：340.至多包含 K 个不同字符的最长子串
// 难度：中等
// 思路：双指针，滑动窗口
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	ans := 0
	hash := map[uint8]int{}
	for left, right := 0, 0; right < len(s); right++ {
		hash[s[right]]++
		if len(hash) == k {
			ans = base.Max(ans, right-left+1)
		}
		for len(hash) > k {
			hash[s[left]]--
			if hash[s[left]] <= 0 {
				delete(hash, s[left])
			}
			left++
		}
	}
	return ans
}
