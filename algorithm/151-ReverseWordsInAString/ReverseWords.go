// https://leetcode-cn.com/problems/reverse-words-in-a-string/
package ReverseWords

import (
	"strings"
)

func tokenize(str string) []string {
	ans := []string{}
	tokens := strings.Split(str, " ")
	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			ans = append(ans, token)
		}
	}
	return ans
}

func reverse(str string) string {
	data := []rune(str)
	for left, right := 0, len(data)-1; left < right; left, right = left+1, right-1 {
		data[left], data[right] = data[right], data[left]
	}
	return string(data)
}

func reverseWords(str string) string {
	reversed := reverse(str)
	tokens := tokenize(reversed)
	for i, token := range tokens {
		tokens[i] = reverse(token)
	}
	return strings.Join(tokens, " ")
}
