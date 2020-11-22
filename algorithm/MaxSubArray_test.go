package algorithm

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := [][]interface{}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).(int)
		if got := maxSubArray2(in); !reflect.DeepEqual(got, want) {
			t.Errorf("maxSubArray(%v).got:%v want:%v", in, got, want)
		}
	}
}