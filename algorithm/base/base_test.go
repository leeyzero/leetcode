package base

import (
	"reflect"
	"testing"
)

func TestUnmarshalTreeNodeNull(t *testing.T) {
	content := "[-1]"
	root, err := UnmarshalTreeNode(content)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", content, err)
	}

	if root != nil {
		t.Errorf("UnmarshalTreeNode(%v) result unexpect.got:%v want:nil", content, root)
	}

}

func TestUnmarshalTreeNodeNormal(t *testing.T) {
	content := "[3,9,20,-1,-1,15,7]"
	root, err := UnmarshalTreeNode(content)
	if err != nil {
		t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", content, err)
	}

	want := []int{3, 9, 20, 15, 7}
	if !reflect.DeepEqual(LevelOrderTraverseTreeNode(root), want) {
		t.Errorf("UnmarshalTreeNode(%v) result unexpect.got:%v want:%v", content, root, want)
	}
}
