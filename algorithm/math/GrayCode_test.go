package math

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := [][]interface{}{
		{0, []int{0}},
		{2, []int{0, 1, 3, 2}},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		want := (test[1]).([]int)
		if got := grayCode(n); !reflect.DeepEqual(got, want) {
			t.Errorf("grayCode(%v).got:%v want:%v", n, got, want)
		}
	}
}
