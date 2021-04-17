package dynamicprogramming

import (
	"testing"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := [][]interface{}{
		{"bbbab", 4},
		{"cbbd", 2},
		{"aecda", 3},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).(int)
		if got := longestPalindromeSubseq(s); got != want {
			t.Errorf("longestPalindromeSubseq(%v).got:%v want:%v", s, got, want)
		}
	}
}
