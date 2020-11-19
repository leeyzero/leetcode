package algorithm

import (
	"testing"
)


func TestGetIntersectionNode(t *testing.T) {
	com := makeLinkList([]int{8, 4, 5})
	headA := makeLinkList([]int{4, 1})
	headB := makeLinkList( []int{5, 0, 1})
	forward(headA, 1).Next = com
	forward(headB, 2).Next = com

	if r, want := getIntersectionNode(headA, headB), com; r != want {
		t.Errorf("getIntersectionNode.got:%v want:%d", r.Val, want.Val)
	}
}