package algorithm

// 剑指 Offer 52. 两个链表的第一个公共节点
// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
// 解题思路：先分别计算出L1和L2的链表长度，计算长度差值k，让较长的链表先走k步，然后两链表同时走，直到相交的节点
// 或到达结尾
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLinkListLength(headA)
	lenB := getLinkListLength(headB)
	currA, currB := headA, headB
	if lenA > lenB {
		currA = forward(headA, lenA-lenB)
	} else {
		currB = forward(headB, lenB-lenA)
	}

	for currA != nil && currB != nil && currA != currB {
		currA = currA.Next
		currB = currB.Next
	}

	if currA == nil || currB == nil {
		return nil
	}
	return currA
}

func getLinkListLength(head *ListNode) int {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	return length
}

