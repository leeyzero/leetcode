package algorithm

import (
	"testing"
	"reflect"
)

func TestLevelOrderBottom(t *testing.T) {
	in := []int{3, 9, -1, -1, 20, 15, -1, -1, 7}
	want := [][]int{
		[]int{15, 7},
		[]int{9, 20},
		[]int{3},
	}
	root := makeTreeNode(in)
	if got := levelOrderBottom(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderBottom(%v).got:%v want:%v", in, got, want)
	}
}