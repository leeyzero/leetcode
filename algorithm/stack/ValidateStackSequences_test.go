package stack

import (
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	}
	for _, test := range tests {
		in1 := (test[0]).([]int)
		in2 := (test[1]).([]int)
		want := (test[2]).(bool)
		if got := validateStackSequences(in1, in2); got != want {
			t.Errorf("ValidateStackSequences(%v, %v).got:%v want:%v", in1, in2, got, want)
		}
	}
}
