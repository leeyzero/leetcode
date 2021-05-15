package queue

// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
// 题目：剑指 Offer 59 - II. 队列的最大值
// 难度：中等
// 思路：单调队列
type MaxQueue struct {
	que    []int
	maxQue []int
}

func NewMaxQueue() *MaxQueue {
	return &MaxQueue{
		que:    []int{},
		maxQue: []int{},
	}
}

func (this *MaxQueue) MaxValue() int {
	if this.isEmpty() {
		return -1
	}
	return this.maxQue[0]
}

func (this *MaxQueue) PushBack(value int) {
	this.que = append(this.que, value)
	for len(this.maxQue) > 0 && this.maxQue[len(this.maxQue)-1] < value {
		this.maxQue = this.maxQue[:len(this.maxQue)-1]
	}
	this.maxQue = append(this.maxQue, value)
}

func (this *MaxQueue) PopFront() int {
	if this.isEmpty() {
		return -1
	}

	front := this.que[0]
	this.que = this.que[1:]
	if front == this.maxQue[0] {
		this.maxQue = this.maxQue[1:]
	}
	return front
}

func (this *MaxQueue) isEmpty() bool {
	return len(this.que) <= 0
}
