package twopointer

// https://leetcode-cn.com/problems/minimum-window-substring/
// 题目：76. 最小覆盖子串
// 难度：困难
// 思路：双指针(滑动窗口)，哈希
func minWindow(s string, t string) string {
	hash := map[uint8]int{}
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	count := 0
	minLeft, minRight := 0, len(s)+1
	for left, right := 0, 0; right < len(s); right++ {
		c := s[right]
		if _, ok := hash[c]; !ok {
			continue
		}

		hash[c] = hash[c] - 1
		if hash[c] >= 0 {
			count++
		}

		// 窗口[left,right]中已经包含了t中的全部字母
		// 尝试left右移，在不影响结果的情况下获得最短字符串
		for count >= len(t) {
			if right-left < minRight-minLeft {
				minLeft = left
				minRight = right
			}
			if _, ok := hash[s[left]]; ok {
				hash[s[left]] = hash[s[left]] + 1
				if hash[s[left]] > 0 {
					count--
				}
			}
			left++
		}
	}

	if minRight > len(s) {
		return ""
	}
	return s[minLeft : minRight+1]
}
