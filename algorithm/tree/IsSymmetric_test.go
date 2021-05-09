package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestIsSymmetric(t *testing.T) {
	in := []int{1, 2, 3, -1, -1, 4, -1, -1, 2, 4, -1, -1, 3, -1, -1}
	root := base.MakeTreeNode(in)
	if got, want := isSymmetric(root), true; !got {
		t.Errorf("isSymmetric(%v).got:%v, want:%v", in, got, want)
	}
}
