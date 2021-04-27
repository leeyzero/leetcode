package queue

// https://leetcode-cn.com/problems/design-circular-queue/
// 题目：622. 设计循环队列
// 难度：中等
// 思路：用数组实现循环队列，front指向对头，tail指向队尾，其中tail可以通过front计算得出
// tail = (front + size - 1) % capacity
type MyCircularQueue struct {
	queue []int
	front int
	size  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func NewMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		front: 0,
		size:  0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	pos := (this.front + this.size) % len(this.queue)
	this.queue[pos] = value
	this.size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % len(this.queue)
	this.size--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.getRearIndex()]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.queue)
}

func (this *MyCircularQueue) getRearIndex() int {
	return (this.front + this.size - 1) % len(this.queue)
}
