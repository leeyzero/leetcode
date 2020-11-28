package algorithm

import (
	"container/list"
)

// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
// 解题思路：最直观的解法是，用两个指针维护窗口大小，每次计算窗口中的最大值，时间复杂度为：O(kn)
// 能不能从获取窗口中最大值的时间复杂度从O(k)降低到O(1)呢？
// 一个滑动窗口可以看成是一个队列，当窗口滑动时，处于窗口的第一个数字被删除，同时在窗口的末尾添加一个新的
// 数字。这看起来像是队列的『先进先出』特性，如果能从这个『队列』中找出最大值，就可以解决这个问题了。
func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	if k <= 0 || k > len(nums) {
		return ans
	}

	deque := list.New()
	left, right := 0, 0
	for right < len(nums) {
		// 维护单调递减队列
		for deque.Len() > 0 && nums[(deque.Back().Value).(int)] < nums[right] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(right)
		if deque.Len() > 0 && (deque.Front().Value).(int) < left {
			deque.Remove(deque.Front())
		}

		size := right - left + 1
		if size < k {
			// 窗口还未完全打开
			right++
			continue
		}

		// 加入窗口最大值
		ans = append(ans, nums[(deque.Front().Value).(int)])
		// 滑动窗口
		left, right = left + 1, right + 1
	}
	return ans
}