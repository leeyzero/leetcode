package twopointer

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
// 题目：524. 通过删除字母匹配到字典里最长单词
// 难度：中等
// 思路：双指针
func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for _, str := range dictionary {
		if isSubSequence(str, s) {
			if len(str) > len(ans) || (len(str) == len(ans) && str < ans) {
				ans = str
			}
		}
	}
	return ans
}

// 判断x是否是y的子序列
func isSubSequence(x string, y string) bool {
	i := 0
	for j := 0; j < len(y) && i < len(x); j++ {
		if x[i] == y[j] {
			i++
		}
	}
	return i == len(x)
}
