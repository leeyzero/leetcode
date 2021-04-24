package twopointer

import (
	"reflect"
	"testing"
)

func TestExchange(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
	}
	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).([]int)
		if got := exchange2(in); !reflect.DeepEqual(got, want) {
			t.Errorf("exchange(%v).got:%v want:%v", in, got, want)
		}
	}
}
