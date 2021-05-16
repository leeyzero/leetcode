package strings

import (
	"strings"
)

// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
// 题目：557. 反转字符串中的单词 III
// 难度：简单
// 描述：给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
func reverseWordsIII(s string) string {
	tokens := strings.Split(s, " ")
	for i, token := range tokens {
		tokens[i] = reverseString(token)
	}
	return strings.Join(tokens, " ")
}

func reverseString(s string) string {
	data := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return string(data)
}
