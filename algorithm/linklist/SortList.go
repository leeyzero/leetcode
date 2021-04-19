package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/sort-list/
// 题目：148. 排序链表
// 难度：中等
// 思路：归并排序
func sortList(head *base.ListNode) *base.ListNode {
	return mergeSort(head, nil)
}

func mergeSort(head *base.ListNode, tail *base.ListNode) *base.ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	return mergeTwoList(mergeSort(head, mid), mergeSort(mid, tail))
}

func mergeTwoList(head1 *base.ListNode, head2 *base.ListNode) *base.ListNode {
	sentinel := &base.ListNode{}
	curr, curr1, curr2 := sentinel, head1, head2
	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			curr.Next = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr2 = curr2.Next
		}
		curr = curr.Next
	}
	if curr1 != nil {
		curr.Next = curr1
	}
	if curr2 != nil {
		curr.Next = curr2
	}
	return sentinel.Next
}
