package search

import (
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	want := true
	if got := exist(board, word); got != want {
		t.Errorf("exist(%v, %v).got:%v want:%v", board, word, got, want)
	}
}
