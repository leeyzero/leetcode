package algorithm

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3}, [][]int{
			[]int{1,2,3},
			[]int{1,3,2},
			[]int{2,1,3},
			[]int{2,3,1},
			[]int{3,2,1},
			[]int{3,1,2},
		}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).([][]int)
		if got := permute(p1); !reflect.DeepEqual(got, want) {
			t.Errorf("permute(%v).got:%v want:%v", p1, got, want)
		}
	}
}