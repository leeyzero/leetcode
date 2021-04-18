package math

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).([]int)
		if got := productExceptSelf(nums); !reflect.DeepEqual(got, want) {
			t.Errorf("productExceptSelf(%v).got:%v want:%v", nums, got, want)
		}
	}
}
