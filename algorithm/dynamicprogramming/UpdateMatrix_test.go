package dynamicprogramming

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
	}
	for _, test := range tests {
		matrix := (test[0]).([][]int)
		want := (test[0]).([][]int)
		if got := updateMatrix(matrix); !reflect.DeepEqual(got, want) {
			t.Errorf("updateMatrix(%v).got:%v want:%v", matrix, got, want)
		}
	}
}
