package search

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := [][]interface{}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}
	for _, test := range tests {
		n := (test[0]).(int)
		k := (test[1]).(int)
		want := (test[2]).([][]int)
		if got := combine(n, k); !reflect.DeepEqual(got, want) {
			t.Errorf("combine(%v, %v).got:%v want:%v", n, k, got, want)
		}
	}
}
