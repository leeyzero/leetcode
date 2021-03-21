package twopointer

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestDetectCycle(t *testing.T) {
	in := []int{3, 2, 0, -4}
	head := base.MakeLinkList(in)
	base.Forward(head, 3).Next = head.Next
	entry := detectCycle(head)
	if entry != head.Next {
		t.Errorf("detectCycle(%v).got:%v want:%v", in, entry, head.Next)
	}
}
