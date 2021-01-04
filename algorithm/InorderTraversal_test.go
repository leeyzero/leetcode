package algorithm

import (
	"testing"
	"reflect"
)

func TestInorderTraversal(t *testing.T) {
	in := []int{2, 1, -1, -1, 3, -1, -1}
	root := makeTreeNode(in)
	if got, want := inorderTraversal(root), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("inorderTraversal(%v).got:%v want:%v", in, got, want)
	}
}

