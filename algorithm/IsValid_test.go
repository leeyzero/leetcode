package algorithm

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := [][]interface{}{
		{"()", true},
		{"()[]{}", true},
		{"[)", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, test := range tests {
		p1 := (test[0]).(string)
		want := (test[1]).(bool)
		if got := isValid(p1); got != want {
			t.Errorf("isFvalid(%v).got:%v want:%v", p1, got, want)
		}
	}
}