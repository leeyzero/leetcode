package MergeTowLists

// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for (l1 != nil) && (l2 != nil) {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	for l1 != nil {
		curr.Next = l1
		l1, curr = l1.Next, curr.Next
	}
	for l2 != nil {
		curr.Next = l2
		l2, curr = l2.Next, curr.Next
	}
	return sentinel.Next
}