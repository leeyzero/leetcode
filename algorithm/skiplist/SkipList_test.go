package skiplist

import (
	"testing"
)

// Skiplist skiplist = new Skiplist();
// skiplist.add(1);
// skiplist.add(2);
// skiplist.add(3);
// skiplist.search(0);   // 返回 false
// skiplist.add(4);
// skiplist.search(1);   // 返回 true
// skiplist.erase(0);    // 返回 false，0 不在跳表中
// skiplist.erase(1);    // 返回 true
// skiplist.search(1);   // 返回 false，1 已被擦除
func TestSkipList(t *testing.T) {
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	if got, want := sl.Search(1), true; got != want {
		t.Errorf("Search(%v) failed.got:%v want:%v", 1, got, want)
	}
	sl.Add(4)
	if got, want := sl.Erase(0), false; got != want {
		t.Errorf("Erase(%v) failed.got:%v want:%v", 0, got, want)
	}
	if got, want := sl.Erase(1), true; got != want {
		t.Errorf("Erase(%v) failed.got:%v want:%v", 1, got, want)
	}
	if got, want := sl.Search(1), false; got != want {
		t.Errorf("Search(%v) failed.got:%v want:%v", 1, got, want)
	}
}
