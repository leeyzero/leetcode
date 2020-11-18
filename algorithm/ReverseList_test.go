package algorithm

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	l1 := makeLinkList([]int{1, 2, 3, 4, 5})
	l2 := reverseList(l1)
	if !reflect.DeepEqual([]int{5, 4, 3, 2, 1}, convertToSlice(l2)) {
		t.Errorf("test reverse list fail")
	}
}
