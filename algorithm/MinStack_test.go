package algorithm

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if got, want := minStack.Min(), -3; got != want {
		t.Errorf("minStack.Min().got:%v want:%v", got, want)
	}
	minStack.Pop()
	if got, want := minStack.Top(), 0; got != want {
		t.Errorf("minStack.Top().got:%v want:%v", got, want)
	}
	if got, want := minStack.Min(), -2; got != want {
		t.Errorf("minStack.Min().got:%v want:%v", got, want)
	}
}

func TestMinStackHasSameMin(t *testing.T) {
	ms := Constructor()
	ms.Push(0)
	ms.Push(1)
	ms.Push(0)
	if got, want := ms.Min(), 0; got != want {
		t.Errorf("ms.Min().got:%v want:%v", got, want)
	}
	ms.Pop()
	if got, want := ms.Min(), 0; got != want {
		t.Errorf("ms.Min().got:%v want:%v", got, want)
	}
}
