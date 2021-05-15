package linklist

// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
// 题目：剑指 Offer 35. 复杂链表的复制
// 难度：中等
// 描述：请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
// 思路：
// 方法一：hash表
// 方法二：拼接+拆分
type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}

	hash := map[*RNode]*RNode{}
	for curr := head; curr != nil; curr = curr.Next {
		hash[curr] = &RNode{curr.Val, nil, nil}
	}
	for curr := head; curr != nil; curr = curr.Next {
		hash[curr].Next = hash[curr.Next]
		hash[curr].Random = hash[curr.Random]
	}
	return hash[head]
}
