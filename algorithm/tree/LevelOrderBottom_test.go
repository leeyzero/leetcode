package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestLevelOrderBottom(t *testing.T) {
	in := []int{3, 9, 20, -1, -1, 15, 7}
	want := [][]int{
		{15, 7},
		{9, 20},
		{3},
	}
	root, err := base.UnmarshalTreeNodeByLevelorder(in)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", in, err)
	}
	if got := levelOrderBottom(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderBottom(%v).got:%v want:%v", in, got, want)
	}
}
