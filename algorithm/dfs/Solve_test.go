package dfs

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := [][]interface{}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
	}
	for _, c := range cases {
		p1 := c[0].([][]byte)
		want := c[1].([][]byte)
		solve(p1)
		if got := p1; !reflect.DeepEqual(got, want) {
			t.Errorf("solve(%v) failed.got:%v want:%v", p1, got, want)
		}
	}
}
