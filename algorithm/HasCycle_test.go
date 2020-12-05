package algorithm

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	p1 := []int{3, 2, 0, -4}
	p2 := 1
	head := makeLinkList(p1)
	forward(head, 3).Next = head
	if got, want := hasCycle(head), true; got != want {
		t.Errorf("hasCycle(%v, %v).got:%v want:%v", p1, p2, got, want)
	}
}