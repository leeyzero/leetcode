package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/insertion-sort-list/
// 题目：147. 对链表进行插入排序
// 难度：中等
// 思路：插入排序
func insertionSortList(head *base.ListNode) *base.ListNode {
	if head == nil {
		return nil
	}

	curr := head.Next
	head.Next = nil
	for curr != nil {
		node := curr
		curr = curr.Next
		head = insertToSortedList(head, node)
	}
	return head
}

func insertToSortedList(head *base.ListNode, node *base.ListNode) *base.ListNode {
	sentinel := &base.ListNode{0, head}
	prev, curr := sentinel, head
	for curr != nil {
		if node.Val < curr.Val {
			node.Next = curr
			prev.Next = node
			break
		}
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		node.Next = nil
		prev.Next = node
	}
	return sentinel.Next
}
