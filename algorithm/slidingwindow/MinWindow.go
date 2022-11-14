package slidingwindow

import (
	"math"
)

// https://leetcode-cn.com/problems/minimum-window-substring/
// 题目：76. 最小覆盖子串
// 难度：困难
// 思路：双指针，滑动窗口，哈希
func minWindow(s string, t string) string {
	window := make(map[byte]int) // 记录窗口内满足条件的字符串出现的次数
	need := make(map[byte]int)   // 记录目标字符应该出现的次数
	for i := range t {
		need[t[i]]++
	}

	// 窗口指针以及用于记录窗口中的字符串符合条件的次数
	left, right, hits := 0, 0, 0
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt
	for right < len(s) {
		// 加入窗口的字符
		c := s[right]
		// 扩大窗口
		right++

		// 当前字符满足条件进行更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				hits++
			}
		}

		// 当前窗口已经满足了条件，收缩优化
		for hits >= len(need) {
			// 更新最小覆盖子串
			if right-left < length {
				start, length = left, right-left
			}

			d := s[left]
			// 缩小窗口
			left++

			// 当前字符满足条件进行更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					hits--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
