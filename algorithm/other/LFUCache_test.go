package other

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	if got, want := lfu.Get(1), 1; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 1, got, want)
	}

	lfu.Put(3, 3)
	if got, want := lfu.Get(2), -1; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 2, got, want)
	}
	if got, want := lfu.Get(3), 3; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 3, got, want)
	}

	lfu.Put(4, 4)
	if got, want := lfu.Get(1), -1; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 1, got, want)
	}
	if got, want := lfu.Get(3), 3; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 3, got, want)
	}
	if got, want := lfu.Get(4), 4; got != want {
		t.Errorf("LFUCache.Get(%v) failed.got:%v want:%v", 4, got, want)
	}
}
