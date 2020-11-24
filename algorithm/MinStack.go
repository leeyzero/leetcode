package algorithm

// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
// 解题思路：用一个栈A保存数据，正常实现Push, Pop, Top接口；用一个辅助
// 栈B保存A中非严格降序的元素
// 注意栈中含有相同最小值的情况
type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minData: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) <= 0 || x <= this.Min() {
		this.minData = append(this.minData, x)
	}
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}

	if this.Top() == this.Min() {
		this.minData = this.minData[:len(this.minData)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	if this.isEmpty() {
		return -1
	}
	return this.minData[len(this.minData)-1]
}

func (this *MinStack) isEmpty() bool {
	return len(this.data) <= 0
}
