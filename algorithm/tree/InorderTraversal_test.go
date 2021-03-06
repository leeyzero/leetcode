package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestInorderTraversal(t *testing.T) {
	in := []int{1, 2, 4, -1, -1, 5, -1, -1, 3}
	root := base.UnmarshalTreeNodeByPreorder(in)
	if got, want := inorderTraversal(root), []int{4, 2, 5, 1, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("inorderTraversal(%v).got:%v want:%v", in, got, want)
	}
}
