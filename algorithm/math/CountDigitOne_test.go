package math

import (
	"testing"
)

func TestCountDigitOne(t *testing.T) {
	tests := [][]int{
		{12, 5},
		{13, 6},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		if got := countDigitOne(p1); got != want {
			t.Errorf("countDigitOne(%v).got:%v want:%v", p1, got, want)
		}
	}
}
