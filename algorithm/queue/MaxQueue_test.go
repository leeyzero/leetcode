package queue

import (
	"testing"
)

func TestMaxQueue(t *testing.T) {
	maxQue := NewMaxQueue()
	maxQue.PushBack(1)
	maxQue.PushBack(2)
	if got, want := maxQue.MaxValue(), 2; got != want {
		t.Errorf("maxQue.MaxValue().got:%v want:%v", got, want)
	}
	if got, want := maxQue.PopFront(), 1; got != want {
		t.Errorf("maxQue.PopFront().got:%v want:%v", got, want)
	}
	if got, want := maxQue.MaxValue(), 2; got != want {
		t.Errorf("maxQue.MaxValue().got:%v want:%v", got, want)
	}
}
