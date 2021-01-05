package algorithm

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	// 1.判断是否有环
	for {
		if fast == nil || fast.Next == nil {
			// 无环
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 2.计算环的节点个数
	k := 1
	for curr := fast; curr.Next != fast; curr = curr.Next {
		k++
	}

	// 3.快指针从头开始向前走k步
	fast = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// 4.慢指针从头和快指针一起走，相遇点即为入口点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}