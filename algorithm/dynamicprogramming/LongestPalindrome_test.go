package dynamicprogramming

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := [][]interface{}{
		{"abb", "bb"},
		{"cbbd", "bb"},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).(string)
		if got := longestPalindrome(s); got != want {
			t.Errorf("longestPalindrome(%v).got:%v want:%v", s, got, want)
		}
	}
}
