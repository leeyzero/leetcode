package array

import (
	"reflect"
	"testing"
)

func TestConstructArr(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).([]int)
		if got := constructArr(p1); !reflect.DeepEqual(got, want) {
			t.Errorf("constructArr(%v).got:%v want:%v", p1, got, want)
		}
	}
}
