package algorithm

import (
	"reflect"
	"testing"
)

func TestMirrorTree(t *testing.T) {
	in := []int{4, 2, 1, -1, -1, 3, -1, -1, 7, 6, -1, -1, 9, -1, -1}
	want := []int{4, 7, 9, 6, 2, 3, 1}
	root := makeTreeNode(in)
	if got := preorderTreeNode(mirrorTree(root)); !reflect.DeepEqual(got, want) {
		t.Errorf("mirrorTree(%v).got:%v want:%v", in, got, want)
	}
}
