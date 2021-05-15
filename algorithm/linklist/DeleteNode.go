package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 题目：剑指 Offer 18. 删除链表的节点
// 描述：给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。
// 难度：简单
// 思路：双指针
func deleteNode(head *base.ListNode, val int) *base.ListNode {
	sentinel := &base.ListNode{0, head}
	pre := sentinel
	for curr := head; curr != nil; curr, pre = curr.Next, pre.Next {
		if curr.Val == val {
			pre.Next = curr.Next
			curr = nil
			break
		}
	}
	return sentinel.Next
}
