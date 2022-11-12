package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	r := base.PreOrderTraverseTreeNode(root)
	if !reflect.DeepEqual(preorder, r) {
		t.Errorf("BuildTree failed.got:%v want:%v", r, preorder)
	}
}
