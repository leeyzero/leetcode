package tree

import (
	"reflect"
	"testing"
)

func TestPostorderNTree(t *testing.T) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
			},
			{
				Val: 3,
			},
		},
	}
	want := []int{2, 3, 1}
	if got := postorderNTree(root); !reflect.DeepEqual(got, want) {
		t.Errorf("postorderNTree(%v).got:%v want:%v", root, got, want)
	}
}
