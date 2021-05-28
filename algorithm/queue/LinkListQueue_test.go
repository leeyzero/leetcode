package queue

import (
	"testing"
)

func TestLinkListQueue(t *testing.T) {
	q := NewLinkListQueue()
	q.PushBack(100)
	if got, want := q.Len(), 1; got != want {
		t.Errorf("q.Len().got:%v want:%v", got, want)
	}
	if got, want := q.Front(), 100; got != want {
		t.Errorf("q.Front().got:%v want:%v", got, want)
	}

	q.PushBack(200)
	if got, want := q.Len(), 2; got != want {
		t.Errorf("q.Len().got:%v want:%v", got, want)
	}
	for _, want := range []int{100, 200} {
		if got := q.PopFront(); got != want {
			t.Errorf("q.PopFront().got:%v want:%v", got, want)
		}
	}
	if got, want := q.Len(), 0; got != want {
		t.Errorf("q.Len().got:%v want:%v", got, want)
	}
}
