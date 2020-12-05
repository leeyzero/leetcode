package algorithm

import (
	"reflect"
	"testing"
)

func TestTreeToDoubleList(t *testing.T) {
	tests := [][]interface{}{
		{[]int{4, 2, 1, -1, -1, 3, -1, -1, 5, -1, -1}, []int{1, 2, 3, 4, 5}},
		{[]int{28, -98, -1, -89, -1, -1, 62, -1, 67, -1, -1}, []int{-98, -89, 28, 62, 67}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).([]int)
		root := makeTreeNode(p1)
		head := treeToDoublyList(root);
		got := []int{head.Val}
		for curr := head.Right; curr != head; curr = curr.Right {
			got = append(got, curr.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("treeToDoublyList(%v).got:%v want:%v", p1, got, want)
		}
	}
}