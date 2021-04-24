package dynamicprogramming

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := [][]interface{}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range tests {
		p1 := (test[0]).(string)
		want := (test[1]).(int)
		if got := lengthOfLongestSubstring(p1); got != want {
			t.Errorf("lengthOfLongestSubstring(%v).got:%v want:%v", p1, got, want)
		}
	}
}
