package DeleteNode

import (
	"testing"
	"reflect"
)

func makeLinkList(vals []int) *ListNode{
	sentinel := &ListNode{}
	curr := sentinel
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return sentinel.Next
}

func convertToSlice(head *ListNode) []int {
	ans := []int{}
	for curr := head; curr != nil; curr = curr.Next {
		ans = append(ans, curr.Val)
	}
	return ans
}

func TestDeleteNode(t *testing.T) {
	head := makeLinkList([]int{4, 5, 1, 9})
	h1 := deleteNode(head, 5)
	want := []int{4, 1, 9}
	if r := convertToSlice(h1); !reflect.DeepEqual(r, want) {
		t.Errorf("test deleteNode.got:%v want:%v", r, want)
	}
}