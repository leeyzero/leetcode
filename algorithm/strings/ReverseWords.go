package strings

import (
	"strings"
)

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
// 题目：剑指 Offer 58 - I. 翻转单词顺序
// 难度：简单
// 描述：输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
func reverseWords(s string) string {
	s = reverse(s)
	tokens := tokenize(s)
	for i, token := range tokens {
		tokens[i] = reverse(token)
	}
	return strings.Join(tokens, " ")
}

func reverse(str string) string {
	data := []byte(str)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return string(data)
}

func tokenize(s string) []string {
	ans := []string{}
	tokens := strings.Split(s, " ")
	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			ans = append(ans, token)
		}
	}
	return ans
}
