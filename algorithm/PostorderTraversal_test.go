package algorithm

import (
	"testing"
	"reflect"
)

func TestPostorderTraversal(t *testing.T) {
	in := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7, -1, -1}
	root := makeTreeNode(in)
	if got, want := postorderTraversal(root), []int{4, 5, 2, 6, 7, 3, 1}; !reflect.DeepEqual(got, want) {
		t.Errorf("postorderTraversal(%v).got:%v, want:%v", in, got, want)
	}
}