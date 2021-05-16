package strings

import (
	"testing"
)

func TestReverseWordsIII(t *testing.T) {
	tests := [][]string{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}
	for _, test := range tests {
		if got := reverseWordsIII(test[0]); got != test[1] {
			t.Errorf("reverseWordsIII(%s).got:%v want:%v", test[0], got, test[1])
		}
	}
}
