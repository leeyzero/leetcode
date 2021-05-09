package dynamicprogramming

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := [][]interface{}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, test := range tests {
		text1 := (test[0]).(string)
		text2 := (test[1]).(string)
		want := (test[2]).(int)
		if got := longestCommonSubsequence(text1, text2); got != want {
			t.Errorf("longestCommonSubsequence(%v, %v).got:%v want:%v", text1, text2, got, want)
		}
	}
}
