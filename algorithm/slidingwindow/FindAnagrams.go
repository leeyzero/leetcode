package slidingwindow

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 题目：438. 找到字符串中所有字母异位词
// 难度：中等
/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
- 1 <= s.length, p.length <= 3 * 104
- s 和 p 仅包含小写字母
*/
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}

	left, right, hits := 0, 0, 0
	ans := []int{}
	for right < len(s) {
		c := s[right]
		// 扩大窗口
		right++

		// 满足条件更新窗口
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				hits++
			}
		}

		// 缩小窗口
		for right-left >= len(p) {
			// 更新结果
			if hits == len(need) && (right-left == len(p)) {
				ans = append(ans, left)
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					hits--
				}
				window[d]--
			}
		}
	}
	return ans
}
