package dynamicprogramming

import "testing"

func TestWordBreak(t *testing.T) {
	tests := [][]interface{}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		wordDict := (test[1]).([]string)
		want := (test[2]).(bool)
		if got := wordBreak(s, wordDict); got != want {
			t.Errorf("wordBreak(%v, %v).got%v want:%v", s, wordDict, got, want)
		}
	}
}
