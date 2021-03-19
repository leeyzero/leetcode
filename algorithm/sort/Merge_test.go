package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := [][]interface{}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}
	for _, test := range tests {
		p1 := (test[0]).([][]int)
		want := (test[1]).([][]int)
		if got := merge(p1); !reflect.DeepEqual(got, want) {
			t.Errorf("merge(%v).got:%v want:%v", p1, got, want)
		}
	}
}
