package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestLevelOrder(t *testing.T) {
	in := []int{3, 9, 20, -1, -1, 15, 7}
	root, err := base.UnmarshalTreeNode(in)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", in, err)
	}

	if got, want := levelOrder(root), []int{3, 9, 20, 15, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrder(%v).got:%v want:%v", in, got, want)
	}
}
