package algorithm

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	tests := [][]interface{}{
		{"42", 42},
		{"-91283472332", -2147483648},
		{"4193 with words", 4193},
		{"3.14159", 3},
	}
	for _, test := range tests {
		in := (test[0]).(string)
		want := (test[1]).(int)
		if got := strToInt(in); got != want {
			t.Errorf("strToInt(%v).got:%v want:%v", in, got, want)
		}
	}
}
