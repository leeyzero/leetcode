package twopointer

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestHasCycle(t *testing.T) {
	p1 := []int{3, 2, 0, -4}
	p2 := 1
	head := base.MakeLinkList(p1)
	base.Forward(head, 3).Next = head
	if got, want := hasCycle(head), true; got != want {
		t.Errorf("hasCycle(%v, %v).got:%v want:%v", p1, p2, got, want)
	}
}
