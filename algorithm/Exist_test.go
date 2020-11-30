package algorithm

import (
	"testing"
)

func TestExist(t *testing.T) {
	board1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	board2 := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	tests := [][]interface{}{
		{board1, "ABCCED", true},
		{board2, "abcd", false},
	}
	for _, test := range tests {
		p1 := (test[0]).([][]byte)
		p2 := (test[1]).(string)
		want := (test[2]).(bool)
		if got := exist(p1, p2); got != want {
			t.Errorf("exist(%v,%v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
