package slidingwindow

import (
	"testing"
)

func TestlengthOfLongestSubstring(t *testing.T) {
	cases := [][]interface{}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 1},
	}
	for _, c := range cases {
		p1 := c[0].(string)
		want := c[1].(int)
		if got := lengthOfLongestSubstring(p1); got != want {
			t.Errorf("lengthOfLongestSubstring(%v) failed.got:%v want:%v", p1, got, want)
		}
	}
}
