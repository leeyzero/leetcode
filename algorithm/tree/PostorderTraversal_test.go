package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestPostorderTraversal(t *testing.T) {
	in := []int{1, 2, 3}
	root, err := base.UnmarshalTreeNodeByLevelorder(in)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", in, err)
	}
	if got, want := postorderTraversal(root), []int{2, 3, 1}; !reflect.DeepEqual(got, want) {
		t.Errorf("postorderTraversal(%v).got:%v, want:%v", in, got, want)
	}
}
