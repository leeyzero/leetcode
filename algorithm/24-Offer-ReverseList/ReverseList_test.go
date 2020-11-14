package ReverseList

import (
	"testing"
	"reflect"
)

func makeLinkList(vals []int) *ListNode {
	fakeHead := &ListNode{0, nil}
	curr := fakeHead
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return fakeHead.Next
}

func convertToSlice(head *ListNode) []int {
	curr := head 
	result := []int{}
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	l1 := makeLinkList([]int{1, 2, 3, 4, 5})
	l2 := reverseList(l1)
	if !reflect.DeepEqual([]int{5, 4, 3, 2, 1}, convertToSlice(l2)) {
		t.Errorf("test reverse list fail")
	}
}