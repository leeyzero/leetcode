package ReverseList

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
type ListNode struct {
	Val int 
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	sentinel := &ListNode{0, nil}
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = sentinel.Next
		sentinel.Next = curr
		curr = nextNode
	}
	return sentinel.Next
}

