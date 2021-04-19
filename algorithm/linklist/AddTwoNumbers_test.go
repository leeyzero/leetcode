package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := [][]interface{}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	}

	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).([]int)
		want := (test[2]).([]int)
		if got := base.LinkListToSlice(addTwoNumbers(base.MakeLinkList(p1), base.MakeLinkList(p2))); !reflect.DeepEqual(got, want) {
			t.Errorf("addTwoNumbers(%v,%v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
