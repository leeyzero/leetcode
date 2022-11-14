package slidingwindow

// https://leetcode.cn/problems/permutation-in-string/
// 题目：567. 字符串的排列
// 难度：中等

/*
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false

提示：
1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母

思路：双指针，滑动窗口，hash
*/
func checkInclusion(s1 string, s2 string) bool {
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}

	hits := 0 // 记录满足条件的字符次数
	for left, right := 0, 0; right < len(s2); {
		c := s2[right]
		// 扩大窗口
		right++

		// 字符满足条件
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				hits++
			}
		}

		// 缩小窗口，s1的子串长度不能超过其本身
		for right-left >= len(s1) {
			if hits == len(need) {
				return true
			}

			d := s2[left]
			// 缩小窗口
			left++
			// 满足条件，更新窗口
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					hits--
				}
				window[d]--
			}
		}
	}
	return false
}
