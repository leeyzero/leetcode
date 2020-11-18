package algorithm

import (
	"reflect"
	"testing"
)

func TestMergeTowLists(t *testing.T) {
	l1 := makeLinkList([]int{1, 2, 4})
	l2 := makeLinkList([]int{1, 3, 4})
	head := mergeTwoLists(l1, l2)
	if !reflect.DeepEqual([]int{1, 1, 2, 3, 4, 4}, convertToSlice(head)) {
		t.Errorf("test merge tow lists fail.got:%v", convertToSlice(head))
	}
}
