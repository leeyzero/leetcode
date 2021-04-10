package queue

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 题目：剑指 Offer 09. 用两个栈实现队列
// 难度：简单
// 思路：用两个栈stack1, stack2
// AppendTail的时候加入到stack1
// DeleteHead时，先将stack1中的所有元素逐一弹出后push到stack2，此时statck2的顺序即为队列的加入顺序
// 只需从stack2中pop栈顶元素
type CQueue struct {
	stack1 []int
	stack2 []int
	size   int
}

func NewCQueue() CQueue {
	return CQueue{
		stack1: []int{},
		stack2: []int{},
		size:   0,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
	this.size++
}

func (this *CQueue) DeleteHead() int {
	if this.isEmpty() {
		return -1
	}

	// 注意：只有在this.stack2为空时才重新从stack1中同步数据
	if len(this.stack2) <= 0 {
		for len(this.stack1) > 0 {
			top := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, top)
		}
	}

	top := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	this.size--
	return top
}

func (this *CQueue) isEmpty() bool {
	return this.size <= 0
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
