package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestReverseBetween(t *testing.T) {
	p1 := []int{1, 2, 3, 4, 5}
	p2, p3 := 2, 4
	want := []int{1, 4, 3, 2, 5}
	if got := base.LinkListToSlice(reverseBetween(base.MakeLinkList(p1), p2, p3)); !reflect.DeepEqual(got, want) {
		t.Errorf("reverseBetween(%v,%v,%v).got:%v want:%v", p1, p2, p3, got, want)
	}

}
