package linklist

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// 题目：92. 反转链表 II
// 链接：https://leetcode-cn.com/problems/reverse-linked-list-ii/
// 难度：中等
func reverseBetween(head *base.ListNode, left int, right int) *base.ListNode {
	if left > right {
		return nil
	}

	fakeHead := base.ListNode{0, head}
	prev, begin := &fakeHead, head
	for i := 1; i < left && begin != nil; i++ {
		begin = begin.Next
		prev = prev.Next
	}

	curr, next := begin, (*base.ListNode)(nil)
	for i := left; i <= right && curr != nil; i++ {
		next = curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr = next
	}
	if begin != nil {
		begin.Next = next
	}
	return fakeHead.Next
}
