package search

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMinDepth(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, 2},
		{[]int{2, -1, 3, -1, 4, -1, 5, -1, 6}, 5},
	}
	for _, test := range tests {
		numbers := (test[0]).([]int)
		root, err := base.UnmarshalTreeNodeByLevelorder(numbers)
		if err != nil {
			t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", numbers, err)
		}
		want := (test[1]).(int)
		if got := minDepth(root); got != want {
			t.Errorf("minDepth(%v).got:%v want:%v", numbers, got, want)
		}
	}
}

func TestMinDepth2(t *testing.T) {
	tests := [][]interface{}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, 2},
		{[]int{2, -1, 3, -1, 4, -1, 5, -1, 6}, 5},
	}
	for _, test := range tests {
		numbers := (test[0]).([]int)
		root, err := base.UnmarshalTreeNodeByLevelorder(numbers)
		if err != nil {
			t.Fatalf("UnmarshalTreeNodeByLevelorder(%v) fail.err:%v", numbers, err)
		}
		want := (test[1]).(int)
		if got := minDepth2(root); got != want {
			t.Errorf("minDepth(%v).got:%v want:%v", numbers, got, want)
		}
	}
}
