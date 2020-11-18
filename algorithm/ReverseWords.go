package algorithm

import (
	"strings"
)

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
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
