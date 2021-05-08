package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMaxDepth(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, -1}, 2},
	}
	for _, test := range tests {
		arr := (test[0]).([]int)
		want := (test[1]).(int)
		root, _ := base.UnmarshalTreeNodeByLevelorder(arr)
		if got := maxDepth(root); got != want {
			t.Errorf("maxDepth(%v).got:%v want:%v", arr, got, want)
		}
	}
}
