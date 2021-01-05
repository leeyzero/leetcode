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
	// 这里可能比较费解，可以这样理解：
	// 如果链表存在环，链表一共有n个节点，那么从头开始走n步，必定就会回到了环的入口点
	// 如果从头到入口点有m个节点（不包括入口节点），环有k个节点，则有n = m + k
	// fast先走了n步后，slow从头开始和fast同时走，当同时走了m步后，fast相当于走了k+m步回到入口点，
	// slow走了m步也到达入口点，fast和slow 相遇
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}