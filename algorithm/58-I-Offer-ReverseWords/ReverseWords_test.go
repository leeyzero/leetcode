package ReverseWords

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := [][]string{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}
	for _, test := range tests {
		if got := reverseWords(test[0]); got != test[1] {
			t.Errorf("reverseWords(%s).got:%s want:%s", test[0], got, test[1])
		}
	}
}