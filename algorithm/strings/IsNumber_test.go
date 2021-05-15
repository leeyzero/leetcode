package strings

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	tests := [][]interface{}{
		{"+100", true},
		{"5e2", true},
		{"-123", true},
		{"3.1416", true},
		{"-1E-16", true},
		{"0123", true},
		{"12e", false},
		{"1a3.14", false},
		{"1.2.3", false},
		{"+-5", false},
		{"12e+5.4", false},
		{"3.", true},
		{".", false},
		{".1", true},
	}
	for _, test := range tests {
		p1 := (test[0]).(string)
		want := (test[1]).(bool)
		if got := isNumber(p1); got != want {
			t.Errorf("isNumber(%v).got:%v want:%v", p1, got, want)
		}
	}
}
