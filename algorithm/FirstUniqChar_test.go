package algorithm

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	tests := [][]interface{}{
		{"abaccdeff", byte('b')},
		{"", byte(' ')},
	}
	for _, test := range tests {
		p1 := (test[0]).(string)
		want := (test[1]).(byte)
		if got := firstUniqChar(p1); got != want {
			t.Errorf("firstUniqChar(%v).got:%v want:%v", p1, got, want)
		}
	}
}
