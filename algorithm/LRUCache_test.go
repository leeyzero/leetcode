package algorithm

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	if r := lru.Get(1); r != 1 {
		t.Errorf("Get(1) fail.got:%d, expect:%d", r, 1)
	}

	lru.Put(3, 3)
	if r := lru.Get(2); r != -1 {
		t.Errorf("Get(2) fail.got:%d, expect:%d", r, -1)
	}
	lru.Put(4, 4)
	if r := lru.Get(1); r != -1 {
		t.Errorf("Get(1) fail.got:%d, expect:%d", r, -1)
	}
	if r := lru.Get(3); r != 3 {
		t.Errorf("Get(3) fail.got:%d, expect:%d", r, 3)
	}
	if r := lru.Get(4); r != 4 {
		t.Errorf("Get(4) fail.got:%d, expect:%d", r, 4)
	}
}
