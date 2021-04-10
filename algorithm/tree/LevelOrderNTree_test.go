package tree

import (
	"reflect"
	"testing"
)

func TestLevelOrderNTree(t *testing.T) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	want := [][]int{
		{1},
		{3, 2, 4},
		{5, 6},
	}
	if got := levelOrderNTree(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderNTree(%v).got:%v want:%v", root, got, want)
	}
}
