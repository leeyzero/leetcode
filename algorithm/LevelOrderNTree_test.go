package algorithm

import (
	"reflect"
	"testing"
)

func TestLevelOrderNTree(t *testing.T) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}
	want := [][]int{
		[]int{1},
		[]int{3, 2, 4},
		[]int{5, 6},
	}
	if got := levelOrderNTree(root); !reflect.DeepEqual(got, want) {
		t.Errorf("levelOrderNTree(%v).got:%v want:%v", root, got, want)
	}
}