package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMergeTowLists(t *testing.T) {
	l1 := base.MakeLinkList([]int{1, 2, 4})
	l2 := base.MakeLinkList([]int{1, 3, 4})
	head := mergeTwoLists(l1, l2)
	if !reflect.DeepEqual([]int{1, 1, 2, 3, 4, 4}, base.LinkListToSlice(head)) {
		t.Errorf("test merge tow lists fail.got:%v", base.LinkListToSlice(head))
	}
}
