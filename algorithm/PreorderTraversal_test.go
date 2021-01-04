package algorithm

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	in := []int{2, 1, -1, -1, 3, -1, -1}
	root := makeTreeNode(in)
	if got, want := preorderTraversal(root), []int{2, 1, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("preorderTraversal(%v).got:%v want:%v", in, got, want)
	}
}