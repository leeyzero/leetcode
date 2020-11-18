package algorithm

import (
	"reflect"
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {
	head := MakeLinkList([]int{1, 2, 3, 4, 5})
	h1 := getKthFromEnd(head, 2)
	if !reflect.DeepEqual([]int{4, 5}, ConvertToSlice(h1)) {
		t.Errorf("test getKthFromEnd fail.got:%v", ConvertToSlice(h1))
	}
}
