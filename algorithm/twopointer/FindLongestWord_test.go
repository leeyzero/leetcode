package twopointer

import (
	"testing"
)

func TestFindLongestWord(t *testing.T) {
	tests := [][]interface{}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		dict := (test[1]).([]string)
		want := (test[2]).(string)
		if got := findLongestWord(s, dict); got != want {
			t.Errorf("findLongestWord(%v, %v).got:%v want:%v", s, dict, got, want)
		}
	}
}
