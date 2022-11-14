package slidingwindow

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串
// 难度：中等
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:l
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

思路：滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int) // 记录窗口中字符的次数
	ans := 0
	for left, right := 0, 0; right < len(s); {
		c := s[right]
		// 扩大窗口
		right++

		// 更新窗口内容
		window[c]++

		// 收缩窗口
		for window[c] > 1 {
			// 更新窗口内容
			window[s[left]]--
			left++
		}

		// 更新最优结果
		ans = base.Max(ans, right-left)
	}
	return ans
}
