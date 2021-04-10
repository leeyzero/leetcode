package twopointer

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// 题目：142. 环形链表 II
// 难度：中等
// 思路：快慢指针
func detectCycle(head *base.ListNode) *base.ListNode {
	// 1.判断是否有环
	slow, fast := head, head
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

	// 3.快指针从头开始先向前走k步
	fast = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// 4.慢指针从头和快指针一起走，相遇点即为入口点
	// 这里可能比较费解，可以这样理解：
	// 如果链表存在环，链表一共有n个节点，那么从头开始走n步，必定就会回到了环的入口点
	// 如果从头到入口点有m个节点（不包括入口节点），环有k个节点，则有n = m + k
	// fast从头开始先走了k步后，slow从头开始和fast同时走，当走了m步后，fast相当于走了n=k+m步回到入口点，
	// slow走了m步也到达入口点，fast和slow 相遇
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 方法二
// 先用快慢指针找到相遇点，此时如果慢指针走了k步，那么快指针就走了2k步，k为环的整数倍（相遇时快指针已经走了n圈）
// 设环入口到相遇点为m，则头到相遇点为k-m，而相遇点再走k-m步也到达入口点
// 此时，只需要将快、慢任一指针指向头，然后同时走k-m步就能在入口点相遇
func detectCycle2(head *base.ListNode) *base.ListNode {
	// 1.判断是否有环
	slow, fast := head, head
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

	// 将快指针指向头，然后快慢指针同时走k-m步走到相遇点即为入口点
	fast = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
