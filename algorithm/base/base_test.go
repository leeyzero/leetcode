package base

import (
	"reflect"
	"testing"
)

func TestUnmarshalTreeNodeNull(t *testing.T) {
	numbers := []int{-1}
	root, err := UnmarshalTreeNode(numbers)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", numbers, err)
	}

	if root != nil {
		t.Errorf("UnmarshalTreeNode(%v) result unexpect.got:%v want:nil", numbers, root)
	}

}

func TestUnmarshalTreeNodeNormal(t *testing.T) {
	numbers := []int{3, 9, 20, -1, -1, 15, 7}
	root, err := UnmarshalTreeNode(numbers)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", numbers, err)
	}

	want := []int{3, 9, 20, 15, 7}
	if !reflect.DeepEqual(LevelOrderTraverseTreeNode(root), want) {
		t.Errorf("UnmarshalTreeNode(%v) result unexpect.got:%v want:%v", numbers, root, want)
	}
}
