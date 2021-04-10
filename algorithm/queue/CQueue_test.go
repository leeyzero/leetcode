package queue

import (
	"testing"
)

func TestCQueue(t *testing.T) {
	cque := NewCQueue()
	cque.AppendTail(1)
	cque.AppendTail(2)
	if got, want := cque.DeleteHead(), 1; got != want {
		t.Errorf("cque.DeleteHead.got:%v want:%v", got, want)
	}
	if got, want := cque.DeleteHead(), 2; got != want {
		t.Errorf("cque.DeleteHead.got:%v want:%v", got, want)
	}
}
