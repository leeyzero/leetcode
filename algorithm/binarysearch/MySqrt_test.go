package binarysearch

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := [][]interface{}{
		{4, 2},
		{8, 2},
	}
	for _, test := range tests {
		x := (test[0]).(int)
		want := (test[1]).(int)
		if got := mySqrt(x); got != want {
			t.Errorf("mySqrt(%v).got:%v want:%v", x, got, want)
		}
	}
}
