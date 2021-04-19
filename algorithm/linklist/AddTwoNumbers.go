package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/add-two-numbers/
// 题目：2. 两数相加
// 难度：中等
// 描述：
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func addTwoNumbers(l1 *base.ListNode, l2 *base.ListNode) *base.ListNode {
	sentinel := base.ListNode{0, nil}
	curr, curr1, curr2 := &sentinel, l1, l2
	carry := 0
	for curr1 != nil || curr2 != nil || carry > 0 {
		val := carry
		if curr1 != nil {
			val += curr1.Val
			curr1 = curr1.Next
		}
		if curr2 != nil {
			val += curr2.Val
			curr2 = curr2.Next
		}
		carry = val / 10
		curr.Next = &base.ListNode{val % 10, nil}
		curr = curr.Next
	}
	return sentinel.Next
}
