package dynamicprogramming

import "testing"

func TestMaximalSquare(t *testing.T) {
	tests := [][]interface{}{
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
		{[][]byte{{'0', '1'}, {'1', '0'}}, 1},
	}
	for _, test := range tests {
		matrix := (test[0]).([][]byte)
		want := (test[1]).(int)
		if got := maximalSquare(matrix); got != want {
			t.Errorf("maximalSquare(%v).got:%v want:%v", matrix, got, want)
		}
	}
}
