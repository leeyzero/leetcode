package algorithm

import (
	"testing"
)

func TestVerifyPostorder(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 6, 3, 2, 5}, false},
		{[]int{1, 3, 2, 6, 5}, true},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(bool)
		if got := verifyPostorder(in); got != want {
			t.Errorf("verifyPostorder(%v).got:%v want:%v", in, got, want)
		}
	}
}