package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestKthSmallest(t *testing.T) {
	p1 := []int{3, 1, 4, -1, 2}
	root, err := base.UnmarshalTreeNodeByLevelorder(p1)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", p1, err)
	}

	want := 1
	k := 1
	if got := kthSmallest(root, k); got != want {
		t.Errorf("kthSmallest(%v, %v).got:%v, want:%v", p1, k, got, want)
	}
}
