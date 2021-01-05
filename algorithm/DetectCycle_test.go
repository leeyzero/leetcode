package algorithm

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	in := []int{3, 2, 0, -4}
	head := makeLinkList(in)
	forward(head, 3).Next = head.Next
	entry := detectCycle(head)
	if entry != head.Next {
		t.Errorf("detectCycle(%v).got:%v want:%v", in, entry, head.Next)
	}
}
