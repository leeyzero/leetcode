package sort

import (
	"container/heap"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 题目：347. 前 K 个高频元素
// 难度：中等
// 要求：算法时间复杂度必须优于O(N*logN)
// 思路：桶排序、最大堆
func topKFrequent(nums []int, k int) []int {
	// 先用一个hash统计每个数字出现的频率
	hash := map[int]int{}
	maxFrequent := 0
	for _, n := range nums {
		hash[n]++
		maxFrequent = base.Max(maxFrequent, hash[n])
	}

	// 使用频次桶存储数字
	buckets := make([][]int, maxFrequent+1)
	for key, cnt := range hash {
		buckets[cnt] = append(buckets[cnt], key)
	}

	// 对桶从后往前遍历，找到k个高频元素
	ans := []int{}
	for i := len(buckets) - 1; i >= 0 && len(ans) < k; i-- {
		for _, n := range buckets[i] {
			ans = append(ans, n)
			if len(ans) >= k {
				break
			}
		}
	}
	return ans
}

// 使用最大堆
func topKFrequent2(nums []int, k int) []int {
	hash := map[int]int{}
	for _, n := range nums {
		hash[n]++
	}

	h := &TopKHeap{}
	heap.Init(h)
	for key, val := range hash {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type TopKHeap [][2]int

// sort 需要实现Len Less Swap接口
func (t TopKHeap) Len() int {
	return len(t)
}

func (t TopKHeap) Less(i, j int) bool {
	return t[i][1] < t[j][1]
}

func (t TopKHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// heap 需要实现Push和Pop接口
func (t *TopKHeap) Push(x interface{}) {
	*t = append(*t, x.([2]int))
}

func (t *TopKHeap) Pop() interface{} {
	old := *t
	x := old[t.Len()-1]
	*t = old[:t.Len()-1]
	return x
}
