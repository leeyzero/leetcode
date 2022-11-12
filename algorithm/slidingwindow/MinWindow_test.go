package slidingwindow

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := [][]string{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
	}
	for _, test := range tests {
		p1 := test[0]
		p2 := test[1]
		want := test[2]
		if got := minWindow(p1, p2); got != want {
			t.Errorf("minWindow(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
