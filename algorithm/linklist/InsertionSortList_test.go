package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestInsertionSortList(t *testing.T) {
	tests := [][]interface{}{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).([]int)
		head := base.MakeLinkList(p1)
		if got := base.LinkListToSlice(insertionSortList(head)); !reflect.DeepEqual(got, want) {
			t.Errorf("sortList(%v).got:%v want:%v", p1, got, want)
		}
	}
}
