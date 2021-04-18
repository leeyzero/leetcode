package array

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := [][]interface{}{
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {4, 3}}},
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		want := (test[1]).([][]int)
		if got := generateMatrix(n); !reflect.DeepEqual(got, want) {
			t.Errorf("generateMatrix(%v).got:%v want:%v", n, got, want)
		}
	}
}
