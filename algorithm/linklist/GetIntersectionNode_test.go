package linklist

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestGetIntersectionNode(t *testing.T) {
	com := base.MakeLinkList([]int{8, 4, 5})
	headA := base.MakeLinkList([]int{4, 1})
	headB := base.MakeLinkList([]int{5, 0, 1})
	base.Forward(headA, 1).Next = com
	base.Forward(headB, 2).Next = com
	if r, want := getIntersectionNode(headA, headB), com; r != want {
		t.Errorf("getIntersectionNode.got:%v want:%d", r.Val, want.Val)
	}
}
