package search

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).([][]int)
		if got := subsets(nums); !reflect.DeepEqual(got, want) {
			t.Errorf("subsets(%v).got:%v want:%v", nums, got, want)
		}
	}
}
