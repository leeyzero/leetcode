package algorithm

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	head := MakeLinkList([]int{4, 5, 1, 9})
	h1 := deleteNode(head, 5)
	want := []int{4, 1, 9}
	if r := ConvertToSlice(h1); !reflect.DeepEqual(r, want) {
		t.Errorf("test deleteNode.got:%v want:%v", r, want)
	}
}
