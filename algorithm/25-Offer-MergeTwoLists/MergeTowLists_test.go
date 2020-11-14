package MergeTowLists

import (
	"reflect"
	"testing"
)

func makeLinkList(vals []int) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return sentinel.Next
}

func convertToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestMergeTowLists(t *testing.T) {
	l1 := makeLinkList([]int{1, 2, 4})
	l2 := makeLinkList([]int{1, 3, 4})
	head := mergeTwoLists(l1, l2)
	if !reflect.DeepEqual([]int{1, 1, 2, 3, 4, 4}, convertToSlice(head)) {
		t.Errorf("test merge tow lists fail.got:%v", convertToSlice(head))
	}
}