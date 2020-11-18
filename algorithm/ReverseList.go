package algorithm

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
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
