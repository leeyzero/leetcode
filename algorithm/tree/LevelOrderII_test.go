package tree

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestLevelOrderII(t *testing.T) {
	in := []int{3, 9, 20, -1, -1, 15, 7}
	root, err := base.UnmarshalTreeNodeByLevelorder(in)
	if err != nil {
		t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", in, err)
	}
	want := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
	if got := levelOrderII(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderII(%v).got:%v want:%v", in, got, want)
	}
}
