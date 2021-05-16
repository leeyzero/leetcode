package dynamicprogramming

import (
	"testing"
)

func TestTranslateNum(t *testing.T) {
	tests := [][]int{
		{12258, 5},
		{1, 1},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		if got := translateNum(p1); got != want {
			t.Errorf("translateNum(%v).got:%v want:%v", p1, got, want)
		}
	}
}
