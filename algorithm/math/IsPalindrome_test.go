package math

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := [][]interface{}{
		{121, true},
		{-121, false},
		{10, false},
	}
	for _, test := range tests {
		p1 := (test[0]).(int)
		want := (test[1]).(bool)
		if got := isPalindrome(p1); got != want {
			t.Errorf("isPalindrome(%v).got:%v want:%v", p1, got, want)
		}
	}
}
