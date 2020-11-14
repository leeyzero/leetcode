package ReverseWords

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	type CaseTable struct {
		Input  string
		Answer string
	}
	table := []CaseTable{
		{"the sky is blue", "blue is sky the"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	}
	for _, v := range table {
		got := reverseWords(v.Input)
		if got != v.Answer {
			t.Errorf("test case fail.\n\tinput:%s \n\tgot:%s \n\texpect:%s", v.Input, got, v.Answer)
		}
	}
}
