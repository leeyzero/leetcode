package twopointer

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
// 题目：524. 通过删除字母匹配到字典里最长单词
// 难度：中等
// 给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回 dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。
//
// 示例 1：
//
// 输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
// 输出："apple"
//
// 示例 2：
// 输入：s = "abpcplea", dictionary = ["a","b","c"]
// 输出："a"
//
// 思路：双指针
func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for _, str := range dictionary {
		if len(str) < len(ans) || (len(str) == len(ans) && str > ans) {
			continue
		}
		if isSubSequence(str, s) {
			ans = str
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
