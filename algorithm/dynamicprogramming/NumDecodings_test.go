package dynamicprogramming

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := [][]interface{}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"06", 0},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).(int)
		if got := numDecodings(s); got != want {
			t.Errorf("numDecodings(%v).got:%v want:%v", s, got, want)
		}
	}
}
