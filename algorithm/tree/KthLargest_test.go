package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestKthLargest(t *testing.T) {
	p1 := []int{3, 1, -1, 2, -1, -1, 4, -1, -1}
	p2 := 1
	root := base.UnmarshalTreeNodeByPreorder(p1)
	if got, want := kthLargest(root, p2), 4; got != want {
		t.Errorf("kthLargest(%v, %v).got:%v want:%v", p1, p2, got, want)
	}
}
