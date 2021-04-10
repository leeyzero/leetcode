package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestPreorderTraversal(t *testing.T) {
	in := []int{1, 2, 3, -1, -1, 4, 5}
	root, err := base.UnmarshalTreeNode(in)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", in, err)
	}
	if got, want := preorderTraversal(root), []int{1, 2, 3, 4, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("preorderTraversal(%v).got:%v want:%v", in, got, want)
	}
}
