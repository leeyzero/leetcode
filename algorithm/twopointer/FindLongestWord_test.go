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

func TestIsSubSequence(t *testing.T) {
	cases := [][]interface{}{
		{"apple", "abpcplea", true},
		{"monkey", "abpcplea", false},
	}
	for _, c := range cases {
		p1 := (c[0]).(string)
		p2 := (c[1]).(string)
		if got, want := isSubSequence(p1, p2), (c[2]).(bool); got != want {
			t.Errorf("isSubSequence(%v, %v) failed.got:%v want:%v", p1, p2, got, want)
		}
	}
}
