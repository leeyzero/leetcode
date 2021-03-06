package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestReversePrint(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 3, 2}, []int{2, 3, 1}},
	}

	for _, test := range tests {
		in := (test[0]).([]int)
		want := (test[1]).([]int)
		if got := reversePrint(base.MakeLinkList(in)); !reflect.DeepEqual(got, want) {
			t.Errorf("reversePrint(%v).got:%v want:%v", in, got, want)
		}
	}
}
