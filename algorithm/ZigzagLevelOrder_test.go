package algorithm

import (
	"testing"
	"reflect"
)

func TestZigzagLevelOrder(t *testing.T) {
	in := []int{3, 9, -1, -1, 20, 15, -1, -1, 7}
	want := [][]int{
		[]int{3}, 
		[]int{20, 9},
		[]int{15, 7},
	}
	root := makeTreeNode(in)
	if got := zigzagLevelOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("zigzagLevelOrder(%v).got:%v want:%v", in, got, want)
	}
}
