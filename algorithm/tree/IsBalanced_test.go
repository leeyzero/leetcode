package tree

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestIsBalanced(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, true},
		{[]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, false},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(bool)
		root, _ := base.UnmarshalTreeNodeByLevelorder(nums)
		if got := isBalanced(root); got != want {
			t.Errorf("isBalanced(%v).got:%v want:%v", nums, got, want)
		}
	}
}
