package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestReverseList(t *testing.T) {
	l1 := base.MakeLinkList([]int{1, 2, 3, 4, 5})
	l2 := reverseList(l1)
	if !reflect.DeepEqual([]int{5, 4, 3, 2, 1}, base.LinkListToSlice(l2)) {
		t.Errorf("test reverse list fail")
	}
}
