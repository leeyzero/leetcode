package algorithm

import (
	"reflect"
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	tests := [][]interface{}{
		{1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, test := range tests {
		in := (test[0]).(int)
		want := (test[1]).([]int)
		if got := printNumbers(in); !reflect.DeepEqual(got, want) {
			t.Errorf("printNumbers(%v).got:%v want:%v", in, got, want)
		}
	}
}