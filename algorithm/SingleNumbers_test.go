package algorithm

import (
	"reflect"
	"testing"
	"sort"
)

func TestSingleNumbers(t *testing.T) {
	tests := [][]interface{}{
		{[]int{4, 1, 4, 6}, []int{1, 6}},
		{[]int{1, 2, 10, 4, 1, 4, 3, 3}, []int{2, 10}},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).([]int)
		got := singleNumbers(in); 
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("singleNumbers(%v).got:%v want:%v", in, got, want)
		}
	}
}