package algorithm

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
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
