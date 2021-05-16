package hash

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
// 题目：剑指 Offer 50. 第一个只出现一次的字符
// 难度：简单
// 描述：在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
// 思路：hash
func firstUniqChar(s string) byte {
	hash := make([]int, 128)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
