package GetKthFromEnd

import (
	"testing"
	"reflect"
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

func TestGetKthFromEnd(t *testing.T) {
	head := makeLinkList([]int{1, 2, 3, 4, 5})
	h1 := getKthFromEnd(head, 2)
	if !reflect.DeepEqual([]int{4, 5}, convertToSlice(h1)) {
		t.Errorf("test getKthFromEnd fail.got:%v", convertToSlice(h1))
	}
}