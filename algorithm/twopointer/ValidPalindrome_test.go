package twopointer

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := [][]interface{}{
		{"aba", true},
		{"abca", true},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).(bool)
		if got := ValidPalindrome(s); got != want {
			t.Errorf("ValidPalindrome(%v).got:%v. want:%v", s, got, want)
		}
	}
}
