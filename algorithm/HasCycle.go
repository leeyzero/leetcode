package algorithm

// https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	sentinel := ListNode{0, head}
	slow, fast := &sentinel, sentinel.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}