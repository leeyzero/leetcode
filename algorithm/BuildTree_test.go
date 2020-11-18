package algorithm

import (
	"reflect"
	"testing"
)

func PreOrderTraverse(root *TreeNode) []int {
	ans := []int{}
	var PreOrderTraverseAux func(node *TreeNode)
	PreOrderTraverseAux = func(node *TreeNode) {
		if node == nil {
			return
		}

		ans = append(ans, node.Val)
		PreOrderTraverseAux(node.Left)
		PreOrderTraverseAux(node.Right)
	}
	PreOrderTraverseAux(root)
	return ans
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	r := PreOrderTraverse(root)
	if !reflect.DeepEqual(preorder, r) {
		t.Errorf("BuildTree fail.got:%v want:%v", r, preorder)
	}
}
