package algorithm

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	in := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7, -1, -1}
	root := makeTreeNode(in)
	if got, want := preorderTraversal(root), []int{1, 2, 4, 5, 3, 6, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("preorderTraversal(%v).got:%v want:%v", in, got, want)
	}
}