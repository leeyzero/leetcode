package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
// 题目：206. 反转链表
// 难度：简单
// 描述：反转一个单链表。
// 思路：虚拟头结点+头插法
func reverseList(head *base.ListNode) *base.ListNode {
	sentinel := &base.ListNode{
		Val:  0,
		Next: nil,
	}
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = sentinel.Next
		sentinel.Next = curr
		curr = nextNode
	}
	return sentinel.Next
}
