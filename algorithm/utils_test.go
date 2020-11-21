package algorithm

import (
	"reflect"
	"testing"
)

func TestMakeTreeNode(t *testing.T) {
	in := []int{1, 2, -1, -1, 3, -1, -1}
	root := makeTreeNode(in)
	want := []int{1, 2, 3}
	if got := preorderTreeNode(root); !reflect.DeepEqual(got, want) {
		t.Errorf("makeTreeNode(%v).got:%v want:%v", in, got, want)
	}
}