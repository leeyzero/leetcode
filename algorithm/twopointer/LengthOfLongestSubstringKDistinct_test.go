package twopointer

import (
	"testing"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := [][]interface{}{
		{"eceba", 2, 3},
		{"aa", 1, 2},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		k := (test[1]).(int)
		want := (test[2]).(int)
		if got := lengthOfLongestSubstringKDistinct(s, k); got != want {
			t.Errorf("lengthOfLongestSubstringKDistinct(%v, %v).got:%v want:%v", s, k, got, want)
		}
	}
}
