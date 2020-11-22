package algorithm

import (
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	tests := [][]string{
		{"We are happy.", "We%20are%20happy."},
	}
	for _, test := range tests {
		if got := replaceSpace(test[0]); got != test[1] {
			t.Errorf("replaceSpace(%v).got:%v want:%v", test[0], got, test[1])
		}
	}
}