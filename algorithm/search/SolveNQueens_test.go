package search

import (
	"reflect"
	"testing"
)

func TestSolveNQueues(t *testing.T) {
	tests := [][]interface{}{
		{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{1, [][]string{{"Q"}}},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		want := (test[1]).([][]string)
		if got := solveNQueens(n); !reflect.DeepEqual(got, want) {
			t.Errorf("solveNQueens(%v).got:%v want:%v", n, got, want)
		}
	}
}
