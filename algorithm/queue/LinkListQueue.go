package queue

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

type LinkListQueue struct {
	sentinel *base.ListNode
	tail     *base.ListNode
	count    int
}

func NewLinkListQueue() *LinkListQueue {
	q := &LinkListQueue{
		sentinel: &base.ListNode{},
		tail:     nil,
		count:    0,
	}
	q.tail = q.sentinel
	return q
}

// 初始化tail -> sentinel
func (t *LinkListQueue) PushBack(val int) {
	t.tail.Next = &base.ListNode{val, nil}
	t.tail = t.tail.Next
	t.count++
}

// 1. 为空
// 2. 只有一个元素
// 3. 有多个元素
func (t *LinkListQueue) PopFront() int {
	if t.Len() <= 0 {
		return -1
	}

	head := t.sentinel.Next
	if t.Len() == 1 {
		t.sentinel.Next = nil
		t.tail = t.sentinel
	} else {
		t.sentinel.Next = head.Next
	}
	t.count--
	return head.Val
}

func (t *LinkListQueue) Front() int {
	if t.Len() <= 0 {
		return -1
	}

	head := t.sentinel.Next
	return head.Val
}

func (t *LinkListQueue) Len() int {
	return t.count
}
