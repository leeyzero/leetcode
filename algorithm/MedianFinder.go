package algorithm

import (
	"errors"
)

// https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
type MedianFinder struct {
	maxHeap []int // 大顶堆，存储数组前半段
	minHeap []int // 小顶堆，存储数组后半段
}

/** initialize your data structure here. */
func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		minHeap: []int{},
		maxHeap: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	m, n := len(this.maxHeap), len(this.minHeap)
	if m == n {
		// m等于n，将元素添加至maxHeap。先将num插入minHeap，再将minHeap的堆顶元素插入maxHeap
		this.minHeap = pushMinHeap(this.minHeap, num)
		this.minHeap, num, _ = popMinHeap(this.minHeap)
		this.maxHeap = pushMaxHeap(this.maxHeap, num)
	} else {
		// m = n+1, 将元素添加至minHeap. 先将num插入maxHeap, 再将maxHeap的堆顶元素插入minHeap
		this.maxHeap = pushMaxHeap(this.maxHeap, num)
		this.maxHeap, num, _ = popMaxHeap(this.maxHeap)
		this.minHeap = pushMinHeap(this.minHeap, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) <= 0 {
		return float64(0)
	}

	m, n := len(this.maxHeap), len(this.minHeap)
	if m == n {
		return float64(this.maxHeap[0]+this.minHeap[0]) / float64(2)
	}
	return float64(this.maxHeap[0])
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func pushMaxHeap(arr []int, num int) []int {
	arr = append(arr, num)
	c := len(arr) - 1
	p := parent(c)
	for p >= 0 && arr[p] < arr[c] {
		arr[p], arr[c] = arr[c], arr[p]
		c = p
		p = parent(c)
	}
	return arr
}

func popMaxHeap(arr []int) ([]int, int, error) {
	if len(arr) <= 0 {
		return nil, -1, errors.New("heap is empty")
	}

	maxVal := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	p := 0
	c := left(p)
	for c < len(arr) {
		if c < len(arr)-1 && arr[c] < arr[c+1] {
			c++
		}
		if arr[p] < arr[c] {
			arr[p], arr[c] = arr[c], arr[p]
		}
		p = c
		c = left(p)
	}
	return arr, maxVal, nil
}

func pushMinHeap(arr []int, num int) []int {
	arr = append(arr, num)
	c := len(arr) - 1
	p := parent(c)
	for p >= 0 && arr[p] > arr[c] {
		arr[p], arr[c] = arr[c], arr[p]
		c = p
		p = parent(c)
	}
	return arr
}

func popMinHeap(arr []int) ([]int, int, error) {
	if len(arr) <= 0 {
		return nil, -1, errors.New("heap is empty")
	}

	minVal := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	p := 0
	c := left(p)
	for c < len(arr) {
		if c < len(arr)-1 && arr[c] > arr[c+1] {
			c++
		}
		if arr[p] > arr[c] {
			arr[p], arr[c] = arr[c], arr[p]
		}
		p = c
		c = left(p)
	}
	return arr, minVal, nil
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
