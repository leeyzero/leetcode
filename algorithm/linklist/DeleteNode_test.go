package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestDeleteNode(t *testing.T) {
	head := base.MakeLinkList([]int{4, 5, 1, 9})
	h1 := deleteNode(head, 5)
	want := []int{4, 1, 9}
	if r := base.LinkListToSlice(h1); !reflect.DeepEqual(r, want) {
		t.Errorf("test deleteNode.got:%v want:%v", r, want)
	}
}
