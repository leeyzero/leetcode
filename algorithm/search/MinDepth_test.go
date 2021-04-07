package search

import (
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestMinDepth(t *testing.T) {
	tests := [][]interface{}{
		{"[3,9,20,-1,-1,15,7]", 2},
		{"[2,-1,3,-1,4,-1,5,-1,6]", 5},
	}
	for _, test := range tests {
		content := (test[0]).(string)
		root, err := base.UnmarshalTreeNode(content)
		if err != nil {
			t.Fatalf("UnmarshalTreeNode(%v) fail.err:%v", content, err)
		}
		want := (test[1]).(int)
		if got := minDepth(root); got != want {
			t.Errorf("minDepth(%v).got:%v want:%v", content, got, want)
		}
	}
}
